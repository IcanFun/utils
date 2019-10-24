package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	_const "github.com/IcanFun/utils/const"
	"github.com/IcanFun/utils/utils/log"

	"github.com/bitly/go-simplejson"
	"github.com/go-redis/redis"

	"github.com/bwmarrin/snowflake"

	"github.com/shopspring/decimal"
)

func StringArrayIntersection(arr1, arr2 []string) []string {
	arrMap := map[string]bool{}
	result := []string{}

	for _, value := range arr1 {
		arrMap[value] = true
	}

	for _, value := range arr2 {
		if arrMap[value] {
			result = append(result, value)
		}
	}

	return result
}

func StringArrayContains(arr1, arr2 []string) bool {
	arrMap := map[string]bool{}

	for _, value := range arr1 {
		arrMap[value] = true
	}

	for _, value := range arr2 {
		if !arrMap[value] {
			return false
		}
	}

	return true
}

func FileExistsInConfigFolder(filename string) bool {
	if len(filename) == 0 {
		return false
	}

	if _, err := os.Stat(FindConfigFile(filename)); err == nil {
		return true
	}
	return false
}

func RemoveDuplicatesFromStringArray(arr []string) []string {
	result := make([]string, 0, len(arr))
	seen := make(map[string]bool)

	for _, item := range arr {
		if !seen[item] {
			result = append(result, item)
			seen[item] = true
		}
	}

	return result
}

func GetHostnameFromSiteURL(siteURL string) string {
	u, err := url.Parse(siteURL)
	if err != nil {
		return ""
	}

	return u.Hostname()
}

func MapToJson(objmap map[string]string) string {
	if b, err := json.Marshal(objmap); err != nil {
		return ""
	} else {
		return string(b)
	}
}

func UrlEncode(str string) string {
	strs := strings.Split(str, " ")

	for i, s := range strs {
		strs[i] = url.QueryEscape(s)
	}

	return strings.Join(strs, "%20")
}

func StructToMap(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(s)
	json.Unmarshal(j, &m)
	return m
}

func StringExistsInArray(s string, arr []string) bool {
	for _, value := range arr {
		if value == s {
			return true
		}
	}
	return false
}

func Round2(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

func Round2String(f float64, n int) string {
	return fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
}

func NewDecimalFromString(s string) (decimal.Decimal, error) {
	if s == "" {
		return decimal.New(0, 0), nil
	}
	return decimal.NewFromString(s)
}

func CurrentTimestamp() int64 {
	return int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
}

func GenerateOrderId() int64 {
	node, _ := snowflake.NewNode(1)
	id := node.Generate().Int64()
	return id
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	case uint64:
		value.SetUint64(v)
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := NewDecimalFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = NewDecimalFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case uint64:
		amount = decimal.NewFromBigInt(big.NewInt(0).SetUint64(v), 0)
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	case *big.Int:
		amount = decimal.NewFromBigInt(v, 0)
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

func FindDir(dir string) (string, bool) {
	fileName := GetAppPath()
	found := false
	if _, err := os.Stat(fileName + "./" + dir + "/"); err == nil {
		fileName, _ = filepath.Abs(fileName + "./" + dir + "/")
		found = true
	} else if _, err := os.Stat(fileName + "../" + dir + "/"); err == nil {
		fileName, _ = filepath.Abs(fileName + "../" + dir + "/")
		found = true
	} else if _, err := os.Stat(fileName + "../../" + dir + "/"); err == nil {
		fileName, _ = filepath.Abs(fileName + "../../" + dir + "/")
		found = true
	}

	return fileName + "/", found
}

func FindConfigFile(fileName string) string {
	if _, err := os.Stat("./config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("./config/" + fileName)
	} else if _, err := os.Stat("../config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("../config/" + fileName)
	} else if _, err := os.Stat(fileName); err == nil {
		fileName, _ = filepath.Abs(fileName)
	}

	return fileName
}

func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetRedisClient(urls string, poolSize, minIdleConns int, dialTimeout, poolTimeout time.Duration) (redisClient redis.UniversalClient) {

	redisUrl, err := url.Parse(urls)

	if err != nil {
		return nil
	}

	urlValue, _ := url.ParseQuery(redisUrl.RawQuery)

	opt := &redis.UniversalOptions{
		Addrs:        []string{redisUrl.Host},
		Password:     urlValue.Get("password"),
		PoolSize:     poolSize,
		MinIdleConns: minIdleConns,
		DialTimeout:  dialTimeout,
		PoolTimeout:  poolTimeout,
	}

	log.Info("GetRedisClient", "connect to redis：", opt.Addrs)

	cli := redis.NewUniversalClient(opt)

	return cli
}

func GetKlineScoreKey(klineType string, closetime int64) (score int) {

	t := time.Unix(closetime/1000, 0)

	year, month, day, hour, minute := t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute()

	switch klineType {
	case "1m":
		score = year%100*100000000 + month*1000000 + day*10000 + hour*100 + minute
	case "3m":
		score = year%100*100000000 + month*1000000 + day*10000 + hour*100 + minute/3
	case "5m":
		score = year%100*100000000 + month*1000000 + day*10000 + hour*100 + minute/5
	case "15m":
		score = year%100*100000000 + month*1000000 + day*10000 + hour*100 + minute/15
	case "30m":
		score = year%100*100000000 + month*1000000 + day*10000 + hour*100 + minute/30
	case "1h":
		score = year%100*1000000 + month*10000 + day*100 + hour
	case "2h":
		score = year%100*1000000 + month*10000 + day*100 + hour/2
	case "4h":
		score = year%100*1000000 + month*10000 + day*100 + hour/4
	case "6h":
		score = year%100*1000000 + month*10000 + day*100 + hour/6
	case "8h":
		score = year%100*1000000 + month*10000 + day*100 + hour/8
	case "12h":
		score = year%100*1000000 + month*10000 + day*100 + hour/12
	case "1d":
		score = year%100*10000 + month*100 + day
	case "3d":
		score = year%100*10000 + month*100 + day/3
	case "1w":
		_, week := t.ISOWeek()
		score = year%100*100 + week
	case "1M":
		score = year%100*100 + month
	}

	return score
}

func join(args ...interface{}) string {
	s := make([]string, len(args))
	for i, v := range args {
		switch v.(type) {
		case string:
			s[i] = v.(string)
		case int64:
			s[i] = strconv.FormatInt(v.(int64), 10)
		case uint64:
			s[i] = strconv.FormatUint(v.(uint64), 10)
		case float64:
			s[i] = strconv.FormatFloat(v.(float64), 'f', 0, 64)
		case bool:
			if v.(bool) {
				s[i] = "1"
			} else {
				s[i] = "0"
			}
		case *big.Int:
			n := v.(*big.Int)
			if n != nil {
				s[i] = n.String()
			} else {
				s[i] = "0"
			}
		default:
			panic("Invalid type specified for conversion")
		}
	}
	return strings.Join(s, ":")
}

func FormatKey(args ...interface{}) string {
	return join(args...)
}

func NewJSON(data []byte) (j *simplejson.Json, err error) {
	j, err = simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func GetHostFromUris(uris []string, scheme string) (redisAddr []string, password string) {
	redisAddr = make([]string, 0)
	for i := range uris {
		redisUrl, err := url.Parse(uris[i])
		if err != nil {
			log.Error("Unsupported scheme '%s'", uris[i])
			return
		}
		if password == "" {
			urlValue, _ := url.ParseQuery(redisUrl.RawQuery)
			password = urlValue.Get("password")
		}

		switch redisUrl.Scheme {
		case scheme:
			redisAddr = append(redisAddr, redisUrl.Host)
		}
	}
	return
}

func GetHmacCode(s string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}

func GetSign(data map[string]interface{}, secretKey string) string {
	var names []string
	for name := range data {
		if name == "signature" {
			continue
		}
		names = append(names, name)
	}
	sort.Strings(names)

	param := ""
	for _, name := range names {
		var v string
		switch d := data[name].(type) {
		case int64, int:
			v = fmt.Sprintf("%d", d)
		case float64:
			v = fmt.Sprintf("%f", d)
		case string:
			v = d
		default:
			v = fmt.Sprintf("%v", d)
		}
		param += fmt.Sprintf("%s=%s&", name, v)
	}
	param = strings.TrimSuffix(param, "&")
	return GetHmacCode(param, secretKey)
}

func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

//unit 1-年 2-月 3-日
func GetSomeTimeAgoMills(num, unit int) int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	year := t.Year()
	switch unit {
	case 1: //年
		var day time.Duration
		for i := 0; i < num; i++ {
			if (IsLeapYear(year-i) && t.Month() > 2) || (IsLeapYear(year-i-1) && t.Month() < 3) { //今年闰年2月份之后 或 去年闰年3月份之前
				day += 366
			} else {
				day += 365
			}
		}

		return (t.UnixNano() - int64(day*24*time.Hour)) / int64(time.Millisecond)
	case 2: //月(30天)
		return (t.UnixNano() - int64(time.Duration(num)*30*24*time.Hour)) / int64(time.Millisecond)
	case 3: //日
		return (t.UnixNano() - int64(time.Duration(num)*24*time.Hour)) / int64(time.Millisecond)
	}

	return 0
}

func GetIpAddress(r *http.Request) string {
	address := r.Header.Get(_const.HEADER_FORWARDED)

	if len(address) == 0 {
		address = r.Header.Get(_const.HEADER_REAL_IP)
	}

	if len(address) == 0 {
		address, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	return address
}

func MapToStruct(m map[string]string, data interface{})  {
	j,_:=json.Marshal(m)
	json.Unmarshal(j,data)
}