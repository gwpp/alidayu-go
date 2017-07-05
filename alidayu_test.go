//
//
//  alidayu-go
//
//  Created by 甘文鹏 on 2017/7/5.
//  Copyright (c) 2017 甘文鹏. All rights reserved.
//

package alidayu

import (
	"testing"

	"github.com/gwpp/alidayu-go/request"
)

const (
	APP_KEY    = "12345678"
	APP_SECRET = "a8d91bae29c3333d58d2222e346e2cde"
)

func TestSendSMS(t *testing.T) {
	client := NewTopClient(APP_KEY, APP_SECRET)
	req := request.NewAlibabaAliqinFcSmsNumSendRequest()
	req.SmsFreeSignName = "你的签名"
	req.RecNum = "13000000000"
	req.SmsTemplateCode = "SMS_12345678"
	req.SmsParam = `{"code":"1234"}`
	response, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", response)
}

func TestQuerySMS(t *testing.T) {
	client := NewTopClient(APP_KEY, APP_SECRET)

	req := request.NewAlibabaAliqinFcSmsNumQueryRequest()
	req.RecNum = "13000000000"
	req.QueryDate = "20170705"
	req.CurrentPage = "1"
	req.PageSize = "10"

	response, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", response)
}
