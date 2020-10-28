package utils

import (
	"bytes"
	"container/list"
	"crypto/rand"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	mrand "math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	_const "github.com/IcanFun/utils/const"

	"github.com/gin-gonic/gin"

	goi18n "github.com/nicksnyder/go-i18n/i18n"
	"github.com/pborman/uuid"
)

const (
	LOWERCASE_LETTERS = "abcdefghijklmnopqrstuvwxyz"
	UPPERCASE_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMBERS           = "0123456789"
	SYMBOLS           = " !\"\\#$%&'()*+,-./:;<=>?@[]^_`|~"

	USER_AUTH_SERVICE_LDAP = "ldap"
	LDAP_SYNC_TASK_NAME    = "LDAP Syncronization"

	IDSIZE = 10
)

const (
	StatusBadRequest = http.StatusBadRequest
)

type StringInterface map[string]interface{}
type StringMap map[string]string
type StringArray []string
type EncryptStringMap map[string]string

type AppOK struct {
	Status string `json:"status" example:"ok"`
}

type AppError struct {
	Id            string `json:"id"`
	Message       string `json:"message"`               // Message to be display to the end user without debugging information
	DetailedError string `json:"detailed_error"`        // Internal error string to help the developer
	RequestId     string `json:"request_id,omitempty"`  // The RequestId that's also set in the header
	StatusCode    int    `json:"status_code,omitempty"` // The http status code
	Where         string `json:"-"`                     // The function where it happened in the form of Struct.Func
	Code          int    `json:"code"`
	params        map[string]interface{}
}

func (er *AppError) Error() string {
	return er.Where + ": " + er.Message + ", " + er.DetailedError
}

func (er *AppError) Translate(T goi18n.TranslateFunc) {
	if er.params == nil {
		er.Message = T(er.Id)
	} else {
		er.Message = T(er.Id, er.params)
	}
}

func (er *AppError) SystemMessage(T goi18n.TranslateFunc) string {
	if er.params == nil {
		return T(er.Id)
	} else {
		return T(er.Id, er.params)
	}
}

func NewAppError(where string, id string, params map[string]interface{}, details string, status int) *AppError {
	ap := &AppError{}
	ap.Id = id
	ap.params = params
	ap.Message = id
	ap.Where = where
	ap.DetailedError = details
	ap.StatusCode = status
	return ap
}

func (er *AppError) ToJson() string {
	b, err := json.Marshal(er)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func NewLocAppError(where string, id string, params map[string]interface{}, details string) *AppError {
	ap := &AppError{}
	ap.Id = id
	ap.params = params
	ap.Message = id
	ap.Where = where
	ap.DetailedError = details
	ap.StatusCode = 500
	return ap
}

// AppErrorFromJson will decode the input and return an AppError
func AppErrorFromJson(data io.Reader) *AppError {
	str := ""
	bytes, rerr := ioutil.ReadAll(data)
	if rerr != nil {
		str = rerr.Error()
	} else {
		str = string(bytes)
	}

	decoder := json.NewDecoder(strings.NewReader(str))
	var er AppError
	err := decoder.Decode(&er)
	if err == nil {
		return &er
	} else {
		return NewLocAppError("AppErrorFromJson", "model.utils.decode_json.app_error", nil, "body: "+str)
	}
}

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

func NewId() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(26) // removes the '==' padding
	return b.String()
}

func NewRandomString(length int) string {
	var b bytes.Buffer
	str := make([]byte, length+8)
	rand.Read(str)
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(str)
	encoder.Close()
	b.Truncate(length) // removes the '==' padding
	return b.String()
}

func MapFromJson(data io.Reader) map[string]string {
	decoder := json.NewDecoder(data)

	var objmap map[string]string
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]string)
	} else {
		return objmap
	}
}

func ArrayFromJson(data io.Reader) []string {
	decoder := json.NewDecoder(data)

	var objmap []string
	if err := decoder.Decode(&objmap); err != nil {
		return make([]string, 0)
	} else {
		return objmap
	}
}

func ArrayToJson(objmap []string) string {
	if b, err := json.Marshal(objmap); err != nil {
		return ""
	} else {
		return string(b)
	}
}

func IsValidHttpUrl(rawUrl string) bool {
	if strings.Index(rawUrl, "http://") != 0 && strings.Index(rawUrl, "https://") != 0 {
		return false
	}

	if _, err := url.ParseRequestURI(rawUrl); err != nil {
		return false
	}

	return true
}

func StringInterfaceToJson(objmap map[string]interface{}) string {
	if b, err := json.Marshal(objmap); err != nil {
		return ""
	} else {
		return string(b)
	}
}

func GetToday0hMills() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t.UnixNano() / int64(time.Millisecond)
}

var reservedName = []string{
	"signup",
	"login",
	"admin",
	"api",
	"oauth",
}

type Status struct {
	Status string `json:"status"`
}

func StatusFromJson(data io.Reader) *Status {
	decoder := json.NewDecoder(data)

	var status Status
	if err := decoder.Decode(&status); err != nil {
		return nil
	} else {
		return &status
	}
}

func GenerateOrderID() int64 {

	t := time.Now() //2018-07-11 15:07:51.8858085 +0800 CST m=+0.004000001

	s1 := mrand.New(mrand.NewSource(time.Now().UnixNano()))

	//获取当前年月日,时分秒
	y := int64(t.Year())   //年
	m := int64(t.Month())  //月
	d := int64(t.Day())    //日
	h := int64(t.Hour())   //小时
	i := int64(t.Minute()) //分钟
	s := int64(t.Second())
	r := int64(s1.Intn(99999))

	orderID := (y*10000000000+m*100000000+d*1000000+h*10000+i*100+s)*100000 + r
	return orderID

}

func GeneratePayCode() string {

	s1 := mrand.New(mrand.NewSource(time.Now().UnixNano()))

	Paycode := fmt.Sprintf("%d", int64(s1.Intn(999999)))

	return Paycode

}

var baseStr string = "ABCDEFGHJKLMNPQRSTUVWXYZ12345679"
var base []byte = []byte(baseStr)

//1544804416
func Base34(n uint64) []byte {
	quotient := n
	mod := uint64(0)
	l := list.New()
	for quotient != 0 {
		mod = quotient % 32
		quotient = quotient / 32
		l.PushFront(base[int(mod)])
	}
	listLen := l.Len()
	if listLen >= 5 {
		res := make([]byte, 0, listLen)
		for i := l.Front(); i != nil; i = i.Next() {
			res = append(res, i.Value.(byte))
		}
		return res
	} else {
		res := make([]byte, 0, 5)
		for i := 0; i < 5; i++ {
			if i < 5-listLen {
				res = append(res, base[0])
			} else {
				res = append(res, l.Front().Value.(byte))
				l.Remove(l.Front())
			}

		}
		return res
	}
}

var source = "E5FCDG3HQ4BNPJ2RSTUV67MWX89KLYZ"

func CreateCode(user_id int64) string {
	num := user_id
	code := []byte("AAAAA")
	i := 4
	l := int64(len(source))
	for num > 0 {
		mod := num % l
		num = (num - mod) / l
		code[i] = source[mod]
		i--
	}
	return string(code)
}

func DecodeCode(code []byte) int {
	source_string := []byte(source)
	n := 0
	l := len(source)
	for k, v := range code {
		if v == 'A' {
			continue
		}
		x := code[k : k+1]
		p := bytes.Index(source_string, x)
		if p < 0 {
			return 0
		}
		fmt.Println(source_string[p])

		n = l*n + p
	}
	return n
}

//目前是6位纯数字
func IsFundPasswordValid(password string) *AppError {
	if len(password) != 6 {
		return NewAppError("User.IsValid",
			"api.user.is_valid.fund_pwd", nil, "", http.StatusBadRequest,
		)
	}

	if result, _ := regexp.MatchString("\\d+", password); !result {
		return NewAppError("User.IsValid",
			"api.user.is_valid.fund_pwd", nil, "", http.StatusBadRequest,
		)
	}

	return nil
}

func GetProtocol(r *http.Request) string {
	if r.Header.Get(_const.HEADER_FORWARDED_PROTO) == "https" || r.TLS != nil {
		return "https"
	} else {
		return "http"
	}
}

func RenderError(c *gin.Context, err interface{}) {
	t, _ := GetTranslationsAndLocale(c.Request)
	switch e := err.(type) {
	case *AppError:
		err.(*AppError).Translate(t)
	case error:
		err = NewAppError("", e.Error(), nil, e.Error(), http.StatusBadRequest)
		err.(*AppError).Translate(t)
	}
	err.(*AppError).Code = err.(*AppError).StatusCode //兼容code
	c.JSON(200, err)
}

func RenderOK(c *gin.Context, o interface{}) {
	if o == nil {
		o = AppOK{Status: "ok"}
	}
	c.JSON(200, o)
}
