package im

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/IcanFun/utils/utils/log"

	"github.com/spf13/viper"
)

//func init() {
//var config string
//flag.StringVar(&config, "c", "../config.toml", "config file")
//flag.Parse()
//viper.SetConfigFile(config)
//err := viper.ReadInConfig()
//
//if err != nil {
//	panic(err)
//}
//
//if viper.GetBool(`"debug") {
//	fmt.Println("Service RUN on DEBUG mode")
//}
//}

func CreateImAccounts(accid, name, props, icon, token string) error {
	client := &http.Client{}

	str := url.Values{}
	str.Set("accid", accid)
	str.Set("name", name)
	str.Set("props", props)
	str.Set("icon", icon)
	str.Set("token", token)
	req, err := http.NewRequest("POST", "https://api.netease.im/nimserver/user/create.action", strings.NewReader(str.Encode()))
	if err != nil {
		// handle error
		log.Error("CreateImAccounts=>NewRequest err:%+v", err)
		return err
	}

	rand.Seed(time.Now().Unix())
	nonce := strconv.Itoa(rand.Intn(10000))
	Curtime := fmt.Sprint(time.Now().Unix())
	//appSecret := "fc5f223bcb61"
	//
	appKey := viper.GetString("netcase.appKey")
	appSecret := viper.GetString("netcase.appSecret")
	sum := sha1.Sum([]byte(appSecret + nonce + Curtime))
	sum2 := []byte{}
	for _, v := range sum {
		sum2 = append(sum2, v)
	}
	checksum := hex.EncodeToString(sum2)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.Header.Set("AppKey", appKey)
	req.Header.Set("Nonce", nonce)
	req.Header.Set("CurTime", Curtime)
	req.Header.Set("CheckSum", checksum)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Error("CreateImAccounts=>ioutil.ReadAll err:%+v", err)
		return err
	}

	fmt.Println(string(body))
	//log.Info(fmt.Sprintf("result from netcase =>%s", string(body)))
	return nil
}
