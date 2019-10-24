package utils

import (
	i18n2 "github.com/IcanFun/utils/i18n"
	"github.com/IcanFun/utils/middleware"
	"github.com/IcanFun/utils/utils"
	"github.com/IcanFun/utils/utils/log"
)

func InitUtils(s *log.LogSettings, jwtSecret string, api middleware.CheckApiKeyFunc, i18n ...utils.I18n) {
	if err := utils.TranslationsPreInit(); err != nil {
		panic(err)
		return
	}
	log.ConfigZapLog(s)
	middleware.JWTSecret = jwtSecret
	middleware.CheckApiKey = api
	for _, value := range i18n {
		i18n2.AddI18n(value)
	}
	i18n2.WriteI18nFile()
}
