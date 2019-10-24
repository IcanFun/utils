package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	_const "github.com/IcanFun/utils/const"
	"github.com/spf13/viper"

	"github.com/IcanFun/utils/utils/log"

	"github.com/gin-gonic/gin"

	"github.com/IcanFun/utils/i18n"

	"github.com/dgrijalva/jwt-go"
	goi18n "github.com/nicksnyder/go-i18n/i18n"

	"github.com/IcanFun/utils/utils"
)

const CTX = "api_ctx"
const Authority = "api_authority"

type CheckApiKeyFunc func(apiKey string, ctx *gin.Context) (claims CustomClaims, appErr *utils.AppError)

type Context struct {
	CustomClaims  CustomClaims
	TokenString   string
	Params        *ApiParams
	Err           *utils.AppError
	T             goi18n.TranslateFunc
	RequestId     string
	IpAddress     string
	Path          string
	siteURLHeader string
}

var JWTSecret string
var CheckApiKey CheckApiKeyFunc

func ApiHandler(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{
		handleFunc:          h,
		requireCustomClaims: false,
		trustRequester:      false,
		isApi:               true,
	}
}

func ApiCustomClaimsRequired(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{
		handleFunc:          h,
		requireCustomClaims: true,
		trustRequester:      false,
		isApi:               true,
	}
}

func AppHandler(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &handler{
		handleFunc:          h,
		requireCustomClaims: false,
		trustRequester:      false,
		isApi:               false,
	}
}

type handler struct {
	handleFunc          func(*Context, http.ResponseWriter, *http.Request)
	requireCustomClaims bool
	trustRequester      bool
	isApi               bool
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("%v - %v", r.Method, r.URL.Path)

	c := &Context{}
	c.T, _ = utils.GetTranslationsAndLocale(r)
	c.RequestId = utils.NewId()
	c.IpAddress = utils.GetIpAddress(r)
	c.Params = ApiParamsFromRequest(r)

	tokenString := ""

	authHeader := r.Header.Get(_const.HEADER_AUTH)
	if len(authHeader) > 6 && strings.ToUpper(authHeader[0:6]) == _const.HEADER_BEARER {
		tokenString = authHeader[7:]
	} else if len(authHeader) > 5 && strings.ToLower(authHeader[0:5]) == _const.HEADER_TOKEN {
		tokenString = authHeader[6:]
	}

	if len(tokenString) == 0 {
		if cookie, err := r.Cookie(_const.SESSION_COOKIE_TOKEN); err == nil {
			tokenString = cookie.Value

			if h.requireCustomClaims && !h.trustRequester {
				if r.Header.Get(_const.HEADER_REQUESTED_WITH) != _const.HEADER_REQUESTED_WITH_XML {
					c.Err = utils.NewLocAppError("ServeHTTP",
						"api.context.session_expired.app_error", nil,
						"tokenString="+tokenString+" Appears to be a CSRF attempt",
					)
					tokenString = ""
				}
			}
		}
	}

	if len(tokenString) == 0 {
		tokenString = r.URL.Query().Get("access_token")
	}

	c.SetSiteURLHeader(GetProtocol(r) + "://" + r.Host)

	w.Header().Set(_const.HEADER_REQUEST_ID, c.RequestId)

	if !h.isApi {
		w.Header().Set("X-Frame-Options", "SAMEORIGIN")
		w.Header().Set("Content-Security-Policy", "frame-ancestors 'self'")
	} else {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			w.Header().Set("Expires", "0")
		}
	}

	if len(tokenString) != 0 && h.requireCustomClaims {
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(JWTSecret), nil
		})

		if err != nil {
			c.Err = utils.NewLocAppError("ServeHTTP",
				"api.context.jwt_parse.app_error", nil, "err="+err.Error()+",tokenString="+tokenString,
			)
			c.Err.StatusCode = http.StatusUnauthorized
		} else if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			c.CustomClaims = *claims
			c.TokenString = tokenString
		} else {
			c.Err = utils.NewLocAppError("ServeHTTP",
				"api.context.jwt_parse.app_error", nil, "tokenString="+tokenString,
			)
			c.Err.StatusCode = http.StatusUnauthorized
		}
	}

	c.Path = r.URL.Path

	if c.Err == nil && h.requireCustomClaims {
		c.CustomClaimsRequired()
	}

	if c.Err == nil {
		h.handleFunc(c, w, r)
	}

	if c.Err != nil {
		c.Err.Translate(c.T)
		c.Err.RequestId = c.RequestId
		c.LogError(c.Err)
		c.Err.Where = r.URL.Path

		if h.isApi {
			w.WriteHeader(c.Err.StatusCode)
			RenderJson(w, c.Err)
		} else {
			if c.Err.StatusCode == http.StatusUnauthorized {
				http.Redirect(w, r,
					c.GetSiteURLHeader()+"/?redirect="+url.QueryEscape(r.URL.Path), http.StatusTemporaryRedirect,
				)
			} else {
				utils.RenderWebError(c.Err.Message, c.Err.DetailedError, "", c.Err.StatusCode, w, r)
			}
		}
	}
}

func (c *Context) SetSiteURLHeader(url string) {
	c.siteURLHeader = strings.TrimRight(url, "/")
}

func (c *Context) GetSiteURLHeader() string {
	return c.siteURLHeader
}

func (c *Context) RemoveCustomClaimsCookie(w http.ResponseWriter, r *http.Request) {
	sessionCookie := &http.Cookie{
		Name:     _const.SESSION_COOKIE_TOKEN,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	userCookie := &http.Cookie{
		Name:   _const.SESSION_COOKIE_USER,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, sessionCookie)
	http.SetCookie(w, userCookie)
}

func (c *Context) SetInvalidParam(parameter string) {
	c.Err = NewInvalidParamError(parameter)
}

func (c *Context) SetInvalidUrlParam(parameter string) {
	c.Err = NewInvalidUrlParamError(parameter)
}

func NewInvalidParamError(parameter string) *utils.AppError {
	err := utils.NewLocAppError("Context",
		i18n.PARAM_ERROR, map[string]interface{}{"Name": parameter}, parameter,
	)
	err.StatusCode = http.StatusBadRequest
	return err
}

func NewInvalidUrlParamError(parameter string) *utils.AppError {
	err := utils.NewLocAppError("Context",
		"api.context.invalid_body_param.app_error", map[string]interface{}{"Name": parameter}, "",
	)
	err.StatusCode = http.StatusBadRequest
	return err
}

func NewTokenError() *utils.AppError {
	err := utils.NewLocAppError("ServeHTTP",
		i18n.JWT_PARSE_ERROR, nil, "token error",
	)
	err.StatusCode = 600
	return err
}

func (c *Context) IsSystemAdmin() bool {
	return CustomClaimsHasPermissionTo(c.CustomClaims, PERMISSION_MANAGE_SYSTEM)
}

func (c *Context) CustomClaimsRequired() {
	if len(c.CustomClaims.UserId) == 0 {
		c.Err = utils.NewAppError("",
			i18n.SESSION_EXPIRED_ERROR, nil, "UserRequired "+c.TokenString, 600,
		)
		return
	}
}

func (c *Context) RequireUserId() *Context {
	if c.Err != nil {
		return c
	}

	if c.Params.UserId == _const.ME {
		c.Params.UserId = c.CustomClaims.UserId
	}

	if len(c.Params.UserId) < 24 {
		c.SetInvalidUrlParam("user_id")
	}

	return c
}

func (c *Context) RequireService() *Context {
	if c.Err != nil {
		return c
	}

	if len(c.Params.Service) == 0 {
		c.SetInvalidUrlParam("service")
	}

	return c
}

func (c *Context) LogError(err *utils.AppError) {
	if err.Id == "web.check_browser_compatibility.app_error" {
		c.LogDebug(err)
	} else {
		log.Error(utils.TDefault("api.context.log.error"), c.Path, err.Where, err.StatusCode,
			c.RequestId, c.CustomClaims.UserId, c.IpAddress, err.SystemMessage(utils.TDefault), err.DetailedError,
		)
	}
}

func (c *Context) LogDebug(err *utils.AppError) {
	log.Debug(utils.TDefault("api.context.log.error"), c.Path, err.Where, err.StatusCode,
		c.RequestId, c.CustomClaims.UserId, c.IpAddress, err.SystemMessage(utils.TDefault), err.DetailedError,
	)
}

func IsApiCall(r *http.Request) bool {
	return strings.Index(r.URL.Path, "/api/") == 0
}

func (c *Context) SetPermissionError(permission *Permission) {
	c.Err = utils.NewLocAppError("Permissions", "api.context.permissions.app_error", nil, "userId="+c.CustomClaims.UserId+", "+"permission="+permission.Id)
	c.Err.StatusCode = http.StatusForbidden
}

func ApiCustomClaimsRequiredMiddleware(ctx *gin.Context) {
	r := ctx.Request
	c := InitContext(ctx)

	c.Path = r.URL.Path

	if c.Err == nil {
		c.CustomClaimsRequired()
	}

	if c.Err != nil {
		c.Err.Translate(c.T)
		c.Err.RequestId = c.RequestId
		c.LogError(c.Err)
		c.Err.Where = r.URL.Path

		utils.RenderError(ctx, c.Err)
		ctx.Abort()
	}

	ctx.Set(CTX, c)
	ctx.Next()
}

func InitContext(ctx *gin.Context) *Context {
	r := ctx.Request
	w := ctx.Writer

	c := &Context{}
	c.T, _ = utils.GetTranslationsAndLocale(r)
	c.RequestId = utils.NewId()
	c.IpAddress = utils.GetIpAddress(r)
	c.Params = ApiParamsFromGinRequest(ctx)
	c.SetSiteURLHeader(GetProtocol(r) + "://" + r.Host)

	w.Header().Set(_const.HEADER_REQUEST_ID, c.RequestId)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		w.Header().Set("Expires", "0")
	}

	if apiKey := r.Header.Get(_const.HEADER_API_KEY); apiKey != "" {
		//验证apiKey
		if CheckApiKey == nil {
			c.Err = utils.NewLocAppError("ServeHTTP",
				i18n.APIKEY_PARSE_ERROR, nil, "apiKey="+apiKey,
			)
		} else {
			c.CustomClaims, c.Err = CheckApiKey(apiKey, ctx)
		}

		return c
	}

	tokenString := ""

	authHeader := r.Header.Get(_const.HEADER_AUTH)
	if len(authHeader) > 6 && strings.ToUpper(authHeader[0:6]) == _const.HEADER_BEARER {
		tokenString = authHeader[7:]
	} else if len(authHeader) > 5 && strings.ToLower(authHeader[0:5]) == _const.HEADER_TOKEN {
		tokenString = authHeader[6:]
	}

	if len(tokenString) == 0 {
		if cookie, err := r.Cookie(_const.SESSION_COOKIE_TOKEN); err == nil {
			tokenString = cookie.Value

			if r.Header.Get(_const.HEADER_REQUESTED_WITH) != _const.HEADER_REQUESTED_WITH_XML {
				c.Err = utils.NewLocAppError("ServeHTTP",
					i18n.SESSION_EXPIRED_ERROR, nil,
					"tokenString="+tokenString+" Appears to be a CSRF attempt",
				)
				c.Err.StatusCode = 600
				tokenString = ""
			}
		}
	}

	if len(tokenString) == 0 {
		tokenString = r.URL.Query().Get("access_token")
	}

	if len(tokenString) != 0 {
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(viper.GetString("JWTSettings.Secret")), nil
		})

		if err != nil {
			c.Err = utils.NewLocAppError("ServeHTTP",
				i18n.JWT_PARSE_ERROR, nil, "err="+err.Error()+",tokenString="+tokenString,
			)
			c.Err.StatusCode = 600
		} else if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			c.CustomClaims = *claims
			c.TokenString = tokenString
		} else {
			c.Err = utils.NewLocAppError("ServeHTTP",
				i18n.JWT_PARSE_ERROR, nil, "tokenString="+tokenString,
			)
			c.Err.StatusCode = 600
		}
	}

	return c
}

func GetApiContext(c *gin.Context) *Context {
	if ctx, exists := c.Get(CTX); exists {
		return ctx.(*Context)
	}
	return nil
}

var ipRequest = sync.Map{}

func CheckRateLimit(ip, request string) bool {
	current := int(time.Now().Unix())
	currentStr := strconv.Itoa(current)
	//limit  100次
	//timeset 600秒
	//限制600秒最多访问100次
	limit, timeset := GetRateLimitConfig()
	if limit == 0 || timeset == 0 {
		return false
	}
	allowanceStr, timestampStr := LoadAllowance(ip, request)
	allowance, _ := strconv.Atoi(allowanceStr)
	timestamp, _ := strconv.Atoi(timestampStr)
	allowance += int(current-timestamp) * limit / timeset
	if allowance > limit {
		allowance = limit
	}

	if allowance < 1 {
		SaveAllowance(ip, request, "0", currentStr)
		//返回true 代表速率超过,进行错误输出
		return true
	} else {
		allowanceStr = strconv.Itoa(allowance - 1)
		SaveAllowance(ip, request, allowanceStr, currentStr)
		//返回false 代表速率未超过
		return false
	}
}

func LoadAllowance(ip, request string) (allowance, timestamp string) {
	res, ok := ipRequest.Load(ip + "_" + request)
	if !ok {
		currentStr := string(time.Now().Unix())
		defaultLimitInt, _ := GetRateLimitConfig()
		defaultLimitStr := strconv.Itoa(defaultLimitInt)
		allowance, timestamp = defaultLimitStr, currentStr
	} else {
		kv := strings.Split(res.(string), "-")
		allowance, timestamp = kv[0], kv[1]
	}
	return

}

func GetRateLimitConfig() (limit, timeset int) {
	return 0, 0
}

func SaveAllowance(ip, request, allowance, current string) {
	ipRequest.Store(ip+"_"+request, allowance+"-"+current)
}

func RateLimit(ctx *gin.Context) {
	if CheckRateLimit(ctx.ClientIP(), ctx.Request.URL.Path) {
		err := utils.NewLocAppError("context.RateLimit", i18n.REQUEST_RATE_LIMIT, nil,
			fmt.Sprintf("ip = %s path %s", ctx.ClientIP(), ctx.Request.URL.Path))
		utils.RenderError(ctx, err)
		//log.Warn("RateLimit=>ip = %s path %s", ctx.ClientIP(), ctx.Request.URL.Path)
		ctx.Abort()
	} else {
		ctx.Next()
	}
}

func RenderJson(w http.ResponseWriter, o interface{}) {
	if b, err := json.Marshal(o); err != nil {
		w.Write([]byte(""))
	} else {
		w.Write(b)
	}
}
