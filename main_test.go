package utils

import (
	"testing"

	"github.com/IcanFun/utils/utils/log"
	"github.com/spf13/viper"
)

func TestLog(t *testing.T) {
	InitUtils(&log.LogSettings{
		EnableConsole: viper.GetBool("LogSettings.EnableConsole"),
		ConsoleLevel:  viper.GetString("LogSettings.ConsoleLevel"),
		EnableFile:    viper.GetBool("LogSettings.EnableFile"),
		FileLevel:     viper.GetString("LogSettings.FileLevel"),
		FileLocation:  viper.GetString("LogSettings.FileLocation"),
	}, "", nil)
	log.Error("11")
}
