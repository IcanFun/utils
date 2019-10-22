package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-xorm/core"

	"github.com/go-xorm/xorm"
)

type XormLogger struct {
	*xorm.SimpleLogger
}

var ignoreSql = map[string]string{
	"SELECT `Id`, `DealId`, `Price`, `Quantity`, `CreateAt`, `DeleteAt`, `CoinSymbol`, `Currency`, `Amount` FROM `Quotations` WHERE (Currency = ?) AND (CoinSymbol = ? AND DeleteAt = 0) ORDER BY `CreateAt` DESC LIMIT 2": "",
	"SELECT * FROM `Deals` LEFT JOIN `Items` AS `i` ON i.Id = ItemId WHERE (Deals.Status = ? AND ? - Deals.CreateAt > PayExpire*1000) LIMIT 100":                                                                           "",
}

func NewXormLogger() *XormLogger {
	logger := &XormLogger{xorm.NewSimpleLogger(os.Stdout)}
	logger.SetLevel(core.LOG_INFO)
	return logger
}

func (x *XormLogger) Infof(format string, v ...interface{}) {
	if len(v) > 1 {
		sql := v[0].(string)

		if _, ok := ignoreSql[sql]; ok {
			return
		}
		sql = strings.ReplaceAll(sql, "?", "%+v")
		args := make([]interface{}, 1, 2)
		if len(v) == 2 {
			args[0] = v[1]
		} else {
			args = v[1].([]interface{})
			for key, value := range args {
				if s, ok := value.(string); ok {
					args[key] = "\"" + s + "\""
				} else if s, ok := value.(time.Weekday); ok {
					args[key] = int(s)
				}
			}
			args = append(args, v[2])
		}

		x.INFO.Output(2, fmt.Sprintf("[SQL] "+sql+" - took: %v", args...))
	}
}
