package utils

import (
	"github.com/IcanFun/utils/middleware"
	"github.com/IcanFun/utils/utils"
	"github.com/IcanFun/utils/utils/log"
)

func InitUtils(s *log.LogSettings, jwtSecret string, api middleware.CheckApiKeyFunc) {
	if err := utils.TranslationsPreInit(); err != nil {
		panic(err)
		return
	}
	log.ConfigZapLog(s)
	middleware.JWTSecret = jwtSecret
	middleware.CheckApiKey = api
}
