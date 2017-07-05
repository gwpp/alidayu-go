//
//
//  alidayu-go
//
//  Created by 甘文鹏 on 2017/7/5.
//  Copyright (c) 2017 甘文鹏. All rights reserved.
//

package alidayu

import (
	"net/http"
	"net/url"

	"errors"

	"encoding/json"

	"sort"

	"fmt"

	"crypto/md5"

	"time"

	"io/ioutil"

	"strings"

	"github.com/gwpp/alidayu-go/request"
)

type TopClient struct {
	AppKey       string
	TargetAppKey string
	Session      string
	Timestamp    string
	Format       string
	V            string
	PartnerId    string
	Simplify     bool

	gatewayUrl string
	secretKey  string
}

func NewTopClient(appKey string, secretKey string) *TopClient {
	top := new(TopClient)
	top.AppKey = appKey
	top.secretKey = secretKey
	top.gatewayUrl = "http://gw.api.taobao.com/router/rest"
	top.Format = "json"
	top.V = "2.0"

	return top
}

func (client *TopClient) Execute(req request.AlibabaRequest) (result map[string]interface{}, err error) {
	params, err := client.buildParams(req)
	if err != nil {
		return
	}
	response, err := http.DefaultClient.PostForm(client.gatewayUrl, params)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return
	}

	result = make(map[string]interface{})
	err = json.Unmarshal(body, &result)

	return
}

func (client *TopClient) buildParams(req request.AlibabaRequest) (params url.Values, err error) {

	if err = req.ParamsIsValid(); err != nil {
		return
	}
	if len(req.GetMethodName()) == 0 {
		err = errors.New("method is required")
		return
	}

	if len(client.AppKey) == 0 {
		err = errors.New("app_key is required")
		return
	}

	// common params
	paramsCommon := make(map[string]string)
	paramsCommon["method"] = req.GetMethodName()
	paramsCommon["app_key"] = client.AppKey
	paramsCommon["target_app_key"] = client.TargetAppKey
	paramsCommon["sign_method"] = "md5"
	paramsCommon["session"] = client.Session
	paramsCommon["timestamp"] = client.timestamp()
	paramsCommon["format"] = client.Format
	paramsCommon["v"] = client.V
	paramsCommon["partner_id"] = client.PartnerId
	if client.Simplify {
		paramsCommon["simplify"] = "1"
	} else {
		paramsCommon["simplify"] = "0"
	}
	data, err := json.Marshal(req)
	if err != nil {
		return
	}

	// request params
	paramsApi := make(map[string]string)
	if err = json.Unmarshal(data, &paramsApi); err != nil {
		return
	}

	// sign
	paramsCommon["sign"] = client.sign(paramsApi, paramsCommon)

	params = make(url.Values)
	for key, value := range paramsApi {
		params.Set(key, value)
	}
	for key, value := range paramsCommon {
		params.Set(key, value)
	}

	return
}

func (client *TopClient) timestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (client *TopClient) sign(paramsApi map[string]string, paramsCommon map[string]string) string {
	ks := make([]string, 0)
	params := make(map[string]string)
	for key, value := range paramsApi {
		ks = append(ks, key)
		params[key] = value
	}

	for key, value := range paramsCommon {
		ks = append(ks, key)
		params[key] = value
	}
	sort.Strings(ks)
	str := ""
	for _, k := range ks {
		str = fmt.Sprintf("%v%v%v", str, k, params[k])
	}

	str = fmt.Sprintf("%v%v%v", client.secretKey, str, client.secretKey)

	hash := md5.Sum([]byte(str))
	sign := fmt.Sprintf("%x", hash)

	return strings.ToUpper(sign)
}
