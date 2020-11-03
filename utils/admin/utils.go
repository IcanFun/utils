package admin

import (
	"strings"

	"github.com/unknwon/com"
)

func String2Type(src string, ty DatabaseType) interface{} {
	switch ty {
	case Int:
		return com.StrTo(src).MustInt64()
	case Float:
		return com.StrTo(src).MustFloat64()
	default:
		return src
	}
}

func SnakeString(s string) string {
	s = strings.ReplaceAll(s, "ID", "Id")
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}
