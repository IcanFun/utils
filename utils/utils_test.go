package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestRound2(t *testing.T) {
	assert.Equal(t, Round2(0.225, 2), 0.23)
}

func TestHmacCode(t *testing.T) {
	key := "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j"
	sign1 := GetHmacCode("symbol=LTCBTC&side=BUY&type=LIMIT&timeInForce=GTC&quantity=1&price=0.1&recvWindow=5000&timestamp=1499827319559", key)
	sign2 := GetSign(map[string]interface{}{
		"symbol":      "LTCBTC",
		"side":        "BUY",
		"type":        "LIMIT",
		"timeInForce": "GTC",
		"quantity":    1,
		"price":       0.1,
		"recvWindow":  5000,
		"timestamp":   "1499827319559",
	}, key)
	assert.Equal(t, sign1, sign2)
}

func TestName(t *testing.T) {
	req, err := http.NewRequest("POST", "https://api.binance.com/api/v3/order", strings.NewReader("symbol=USDTBTC"))
	if err != nil {
		// handle error
		t.Error("NewRequest", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-MBX-APIKEY", "suIEK2YNWAnV6qFEW94YqCIfcCts8nMLiw9HbBNmvTuOzR2t1wlQntOGw9XH1wSZ")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		t.Error("do ", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		t.Error("ReadAll", err.Error())
		return
	}

	t.Log("result:", string(body))
}

func TestGetAYearAgoMills(t *testing.T) {
	ti, _ := time.ParseInLocation("2006-01-02", "2009-02-03", time.Local)
	year := ti.Year()
	var day time.Duration
	if ((year%4 == 0 && year%100 != 0) || year%400 == 0) && ti.Month() > 2 { //闰年2月份之后
		day = 366
	} else if (((year-1)%4 == 0 && (year-1)%100 != 0) || (year-1)%400 == 0) && ti.Month() < 3 { //去年是闰月
		day = 366
	} else {
		day = 365
	}
	t.Log((ti.UnixNano() - int64(day*24*time.Hour)) / int64(time.Millisecond))
}

func TestIndexString(t *testing.T) {
	t.Log(int(time.Now().Weekday()))
	t.Log([]byte("1111100")[5-1])
}
