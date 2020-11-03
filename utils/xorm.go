package utils

import (
	"fmt"
	"time"

	"github.com/IcanFun/utils/utils/log"
	"xorm.io/core"

	"xorm.io/xorm"
)

type MysqlCons struct {
	DriverName   string
	DataSources  []string
	MaxIdleConns int
	MaxOpenConns int
	NoShowSql    bool
}

func XormConnection(mysql *MysqlCons) *xorm.EngineGroup {
	conns := make([]string, len(mysql.DataSources))
	for key, value := range mysql.DataSources {
		conns[key] = fmt.Sprintf("%s?charset=utf8&parseTime=True&loc=Local", value)
	}
	eg, err := xorm.NewEngineGroup(mysql.DriverName, conns)
	if err != nil {
		panic(err)
	}
	eg.SetMaxIdleConns(mysql.MaxIdleConns)
	eg.SetMaxOpenConns(mysql.MaxOpenConns)
	eg.SetConnMaxLifetime(time.Minute)
	if !mysql.NoShowSql {
		eg.SetLogger(log.NewXormLogger())
		eg.ShowSQL(true)
		eg.SetLogLevel(core.LOG_INFO)
		eg.ShowExecTime(true)
	}
	eg.SetMapper(core.SameMapper{})

	return eg
}
