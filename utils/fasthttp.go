package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/IcanFun/utils/utils/json"

	"github.com/valyala/fasthttp"
)

func Map2Url(obj map[string]interface{}) string {
	var url string
	for key, value := range obj {
		var v string
		switch d := value.(type) {
		case int64, int:
			v = fmt.Sprintf("%d", d)
		case float64:
			v = fmt.Sprintf("%f", d)
		case string:
			v = d
		default:
			v = fmt.Sprintf("%v", d)
		}
		url += fmt.Sprintf("%s=%s&", key, v)
	}
	url = strings.TrimSuffix(url, "&")
	return url
}

func FastGet(URI string, header map[string]string, param map[string]interface{}, res interface{}) (string, error) {
	p := ""
	if param != nil {
		p = Map2Url(param)
		p = "?" + p
	}
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("GET")
	for key, value := range header {
		req.Header.Set(key, value)
	}
	req.SetRequestURI(URI + p)
	client := &fasthttp.Client{}

	err := client.Do(req, resp)
	if err != nil {
		return URI + p, err
	}
	b := resp.Body()
	err = json.Unmarshal(b, res)
	if err != nil {
		err = errors.New(fmt.Sprintf("Unmarshal b[%s] err:%s", string(b), err.Error()))
	}
	return URI + p, err
}

func FastPost(URI string, header map[string]string, param map[string]interface{}, res interface{}) (string, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	for key, value := range header {
		req.Header.Set(key, value)
	}
	req.SetRequestURI(URI)

	requestBody, _ := json.Marshal(param)
	req.SetBody(requestBody)

	err := fasthttp.Do(req, resp)
	if err != nil {
		return URI, err
	}
	b := resp.Body()
	err = json.Unmarshal(b, res)
	if err != nil {
		err = errors.New(fmt.Sprintf("Unmarshal b[%s] err:%s", string(b), err.Error()))
	}
	return URI, err
}
