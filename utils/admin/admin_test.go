package admin

import (
	"testing"

	"github.com/IcanFun/utils/utils"
	"github.com/gin-gonic/gin"
)

type ChangeModel struct {
	ID       int64  `json:"id"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
	DeleteAt *int64 `json:"delete_at,omitempty"`
}

type User struct {
	ChangeModel
	Username string `json:"username"`
	Update   bool   `json:"update"`
}

func (User) TableName() string {
	return "users"
}

type Email struct {
	ID         int
	UserID     int    `gorm:"index"`                          // 外键 (属于), tag `index`是为该列创建索引
	Email      string `gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
	Subscribed bool
}

type UserEmail struct {
	User
	Email
}

func TestAdmin(t *testing.T) {
	db := utils.XormConnection(&utils.MysqlCons{
		DriverName:   "mysql",
		DataSources:  []string{"root:123456@tcp(192.168.31.201:3306)/fg?charset=utf8&parseTime=True&loc=Local"},
		MaxIdleConns: 10,
		MaxOpenConns: 10,
		NoShowSql:    false,
	})

	g := gin.Default()

	admin := InitAdmin(db, g.Group("/admin"))
	option := admin.Table(&UserEmail{}, "ID", "/users").
		SetSelect(&CurdCon{
			Join: []JoinCon{{
				JoinTable: "emails",
				ON:        "emails.user_id = users.id",
			}},
		}).
		SetAdd(&CurdCon{
			Open:    true,
			MwParam: []string{"Username"},
			Mw: []gin.HandlerFunc{func(context *gin.Context) {
				context.Set("Username", "d")
			}},
		}).
		SetEdit(&CurdCon{Open: true}).
		SetDel(&CurdCon{Open: true})

	option.SetFilter(map[string]FilterType{
		"id": {
			FilterOperator: FilterOperatorEqual,
			DatabaseType:   Int,
		},
		"username": {
			FilterOperator: FilterOperatorLike,
			DatabaseType:   String,
		},
		"update": {
			FilterOperator: FilterOperatorIn,
			DatabaseType:   Float,
		},
		"update_at": {
			FilterOperator: FilterOperatorBetween,
			DatabaseType:   Int,
		},
	})
	admin.Start()
	g.Run(":8080")
}
