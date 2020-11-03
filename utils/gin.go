package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type,Authorization")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
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

func RenderErrorV2(c *gin.Context, err error) {
	t, _ := GetTranslationsAndLocale(c.Request)
	appErr, ok := err.(*AppError)
	if !ok {
		appErr = NewAppError("", err.Error(), nil, err.Error(), http.StatusBadRequest)
	}
	appErr.Translate(t)
	if appErr.Code == 0 {
		if appErr.StatusCode == 0 {
			appErr.Code = 400
		} else {
			appErr.Code = appErr.StatusCode
		}
	}
	c.JSON(200, map[string]interface{}{"code": appErr.Code, "date": "", "msg": appErr.Error()})
}

func RenderOKV2(c *gin.Context, o interface{}) {
	if o == nil {
		o = AppOK{Status: "ok"}
	}
	c.JSON(200, map[string]interface{}{"code": 200, "data": o, "msg": ""})
}
