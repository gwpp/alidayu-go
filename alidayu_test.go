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
	//APP_KEY    = "12345678"
	//APP_SECRET = "a8d91bae29c3333d58d2222e346e2cde"

	APP_KEY            = "24534275"
	APP_SECRET         = "a8d91bae29c30f5d58d285be346e2cde"
	PHONE_NUM          = "17706427626"
	SMS_TEMPLATE_CODE  = "SMS_75790179"
	SMS_FREE_SIGN_NAME = "愿望卡"
	CALL_SHOW_NUM      = "4001112222"
)

func TestSendSMS(t *testing.T) {
	client := NewTopClient(APP_KEY, APP_SECRET)
	req := request.NewAlibabaAliqinFcSmsNumSendRequest()
	req.SmsFreeSignName = SMS_FREE_SIGN_NAME
	req.RecNum = PHONE_NUM
	req.SmsTemplateCode = SMS_TEMPLATE_CODE
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
	req.RecNum = PHONE_NUM
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

func TestSingleCall(t *testing.T) {
	client := NewTopClient(APP_KEY, APP_SECRET)

	req := request.NewAlibabaAliqinFcVoiceNumSinglecallRequest()
	req.CalledNum = PHONE_NUM
	req.CalledShowNum = CALL_SHOW_NUM
	req.VoiceCode = "c2e99ebc-2d4c-4e78-8d2a-afbb06cf6216.wav"

	response, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", response)
}
