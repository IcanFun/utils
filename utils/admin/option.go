package admin

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/IcanFun/utils/utils"
	"github.com/IcanFun/utils/utils/log"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"xorm.io/xorm"
)

type FilterType struct {
	FilterOperator
	DatabaseType
}

type table interface {
	TableName() string
}

type JoinCon struct {
	JoinTable  string
	TableAlias string //连表别名
	ON         string //连表条件
}

type CurdCon struct {
	Open    bool              //能否增删改查
	Select  string            //mysql select 的列
	Func    gin.HandlerFunc   //自定义方法
	Mw      []gin.HandlerFunc //增删改查中间件
	MwParam []string          //中间件保存的参数，用于增删改查的参数补充
	Join    []JoinCon
}

type Option struct {
	Table               table  //表格struct
	Key                 string //表主键
	Url                 string //注册路由
	tablePrtType        reflect.Type
	filter              map[string]FilterType //查询所用的筛选条件
	add, del, edit, sel CurdCon
	globalMw            []gin.HandlerFunc //中间件,全局通用的
	globalMwParam       []string          //中间件保存的参数，用于增删改查的参数补充,全局通用的
	renderV             int
}

func GetOpenCurd() *CurdCon {
	return &CurdCon{Open: true}
}

func (c *CurdCon) SetFunc(fun gin.HandlerFunc) *CurdCon {
	c.Func = fun
	return c
}

func (c *CurdCon) SetMw(mws ...gin.HandlerFunc) *CurdCon {
	c.Mw = mws
	return c
}

func (c *CurdCon) SetMwParam(params ...string) *CurdCon {
	c.MwParam = params
	return c
}

func (c *CurdCon) SetJoin(joins ...JoinCon) *CurdCon {
	c.Join = joins
	return c
}

func (o *Option) SetGlobalMw(mws ...gin.HandlerFunc) *Option {
	o.globalMw = mws
	return o
}

func (o *Option) SetGlobalMwParam(keys []string) *Option {
	o.globalMwParam = keys
	return o
}

func (o *Option) SetSelect(con *CurdCon) *Option {
	o.sel = *con
	return o
}

func (o *Option) SetAdd(con *CurdCon) *Option {
	o.add = *con
	return o
}

func (o *Option) SetEdit(con *CurdCon) *Option {
	o.edit = *con
	return o
}

func (o *Option) SetDel(con *CurdCon) *Option {
	o.del = *con
	return o
}

func (o *Option) SetFilter(filter map[string]FilterType) *Option {
	o.filter = filter
	return o
}

func (o *Option) GetSelectFunc(db *xorm.EngineGroup) gin.HandlerFunc {
	if o.sel.Func == nil {
		o.sel.MwParam = append(o.sel.MwParam, o.globalMwParam...)
		if o.sel.Select == "" {
			o.sel.Select = "*"
		}
		o.sel.Func = func(context *gin.Context) {
			limit := com.StrTo(context.Query("limit")).MustInt()
			offset := com.StrTo(context.Query("offset")).MustInt()
			if limit == 0 {
				limit = 10
			}

			var total int64
			list := reflect.New(reflect.SliceOf(o.tablePrtType))
			session := db.Limit(limit, offset)

			for key, operator := range o.filter {
				data := context.Query(key)
				if data == "" {
					continue
				}
				key = strings.ReplaceAll(key, "_", "")
				switch operator.FilterOperator {
				case FilterOperatorLike:
					has := false
					if strings.HasPrefix(key, "%") {
						data = "%" + data
						has = true
					}
					if strings.HasSuffix(key, "%") {
						data = data + "%"
						has = true
					}
					if !has {
						data = "%" + data + "%"
					}
					session = session.Where(fmt.Sprintf("%s like ?", strings.ReplaceAll(key, "%", "")), data)
				case FilterOperatorIn:
					s := strings.Split(data, ",")
					d := make([]interface{}, len(s))
					for key, value := range s {
						d[key] = String2Type(value, operator.DatabaseType)
					}
					session = session.Where(fmt.Sprintf("%s in (?)", key), d)
				case FilterOperatorBetween:
					s := strings.Split(data, "-")
					if len(s) == 2 {
						session = session.Where(fmt.Sprintf("%s between ? AND ?", key), String2Type(s[0], operator.DatabaseType), String2Type(s[1], operator.DatabaseType))
					}
				default:
					session = session.Where(fmt.Sprintf("%s %s ?", key, operator.FilterOperator), String2Type(data, operator.DatabaseType))
				}
			}
			for _, value := range o.sel.MwParam {
				if data, ok := context.Get(value); ok {
					session = session.Where(fmt.Sprintf("%s = ?", value), data)
				}
			}

			for _, value := range o.sel.Join {
				var tablename interface{}
				if value.TableAlias != "" {
					tablename = []string{value.JoinTable, value.TableAlias}
				} else {
					tablename = value.JoinTable
				}
				session = session.Join("LEFT", tablename, value.ON)
			}

			log.Debug("%+v %+v %+v", list, list.Elem(), list.Interface())
			if orderBy := context.Query("order_by"); orderBy != "" {
				session = session.OrderBy(orderBy)
			} else {
				session = session.Desc(o.Table.TableName() + ".Id")
			}

			total, err := session.Select(o.sel.Select).FindAndCount(list.Interface())
			if err != nil {
				log.Error("SelectFunc=>Find error:%s", err.Error())
				o.RenderError(context, err)
				return
			}

			o.RenderOk(context, map[string]interface{}{"list": list.Interface(), "total": total})
		}
	}
	return o.sel.Func
}

func (o *Option) GetAddFunc(db *xorm.EngineGroup) gin.HandlerFunc {
	if o.add.Func == nil {
		o.add.MwParam = append(o.add.MwParam, o.globalMwParam...)

		o.add.Func = func(context *gin.Context) {
			req := reflect.New(o.tablePrtType.Elem())
			if err := context.ShouldBind(req.Interface()); err != nil {
				log.Error("AddFunc=>param error:%s", err.Error())
				o.RenderError(context, err)
				return
			}
			for _, value := range o.add.MwParam {
				if data, ok := context.Get(value); ok {
					log.Info("%+v %+v", req.Elem(), req.Elem().FieldByName(value))
					if req.Elem().FieldByName(value).IsValid() && req.Elem().FieldByName(value).IsZero() {
						req.Elem().FieldByName(value).Set(reflect.ValueOf(data))
					}
				}
			}
			if req.Elem().FieldByName("CreateAt").IsValid() {
				req.Elem().FieldByName("CreateAt").Set(reflect.ValueOf(utils.GetMillis()))
			}
			if req.Elem().FieldByName("UpdateAt").IsValid() {
				req.Elem().FieldByName("UpdateAt").Set(reflect.ValueOf(utils.GetMillis()))
			}
			log.Debug("%+v %+v %+v", req, req.Elem(), req.Interface())
			_, err := db.Insert(req.Interface())
			if err != nil {
				log.Error("AddFunc=>Find error:%s", err.Error())
				o.RenderError(context, err)
				return
			}
			o.RenderOk(context, req.Interface())
		}
	}
	return o.add.Func
}

func (o *Option) GetEditFunc(db *xorm.EngineGroup) gin.HandlerFunc {
	if o.edit.Func == nil {
		o.edit.MwParam = append(o.edit.MwParam, o.globalMwParam...)

		o.edit.Func = func(context *gin.Context) {
			req := map[string]interface{}{}
			if err := context.ShouldBind(&req); err != nil {
				log.Error("EditFunc=>param error:%s", err.Error())
				o.RenderError(context, err)
				return
			}

			key := SnakeString(o.Key)
			if _, ok := req[key]; !ok {
				o.RenderError(context, errors.New(key+"必填"))
				return
			}
			for _, value := range o.edit.MwParam {
				if data, ok := context.Get(value); ok {
					req[value] = data
				}
			}

			if _, ok := req["update_at"]; !ok {
				if _, ok := o.tablePrtType.Elem().FieldByName("UpdateAt"); ok {
					req["UpdateAt"] = utils.GetMillis()
				}
			}

			log.Debug("%+v", req)
			updateValue := map[string]interface{}{}
			for key, value := range req {
				key := strings.ToLower(strings.ReplaceAll(key, "_", ""))
				if _, ok := updateValue[key]; ok {
					log.Error("edit=>req[%v]", req)
					o.RenderError(context, errors.New("非法请求"))
					return
				}
				updateValue[key] = value
			}
			_, err := db.Table(o.Table.TableName()).
				Where(fmt.Sprintf("`%s`.`%s` = ?", o.Table.TableName(), o.Key), updateValue[strings.ReplaceAll(key, "_", "")]).
				Update(updateValue)
			if err != nil {
				log.Error("EditFunc=>Find error:%s", err.Error())
				o.RenderError(context, err)
				return
			}
			o.RenderOk(context, req)
		}
	}
	return o.edit.Func
}

func (o *Option) GetDelFunc(db *xorm.EngineGroup) gin.HandlerFunc {
	if o.del.Func == nil {
		o.del.MwParam = append(o.del.MwParam, o.globalMwParam...)

		o.del.Func = func(context *gin.Context) {
			req := reflect.New(o.tablePrtType.Elem())
			if err := context.ShouldBind(req.Interface()); err != nil {
				log.Error("DelFunc=>param error:%s", err.Error())
				o.RenderError(context, err)
				return
			}

			if req.Elem().FieldByName(o.Key).IsZero() {
				o.RenderError(context, errors.New(o.Key+"必填"))
				return
			}

			session := db.NewSession()
			defer session.Close()
			for _, value := range o.del.MwParam {
				if data, ok := context.Get(value); ok {
					session = session.Where(fmt.Sprintf("%s = ?", value), data)
				}
			}

			log.Debug("%+v %+v %+v %v", req, req.Elem(), req.Interface())
			if req.Elem().FieldByName("DeleteAt").IsValid() {
				_, err := session.Table(o.Table.TableName()).Where(fmt.Sprintf("`%s`.`%s` = ?", o.Table.TableName(), o.Key), req.Elem().FieldByName(o.Key).Interface()).
					Update(map[string]interface{}{"DeleteAt": utils.GetMillis()})
				if err != nil {
					log.Error("EditFunc=>Find error:%s", err.Error())
					o.RenderError(context, err)
					return
				}
			} else {
				_, err := session.Where(fmt.Sprintf("`%s`.`%s` = ?", o.Table.TableName(), o.Key), req.Elem().FieldByName(o.Key).Interface()).Delete(req.Interface())
				if err != nil {
					log.Error("DelFunc=>Find error:%s", err.Error())
					o.RenderError(context, err)
					return
				}
			}

			o.RenderOk(context, nil)
		}
	}
	return o.del.Func
}

func (o *Option) RenderError(c *gin.Context, err error) {
	if o.renderV == 2 {
		utils.RenderErrorV2(c, err)
	} else {
		utils.RenderError(c, err)
	}
}

func (o *Option) RenderOk(c *gin.Context, d interface{}) {
	if o.renderV == 2 {
		utils.RenderOKV2(c, d)
	} else {
		utils.RenderOK(c, d)
	}
}
