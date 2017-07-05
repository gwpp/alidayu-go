//
//
//  alidayu-go
//
//  Created by 甘文鹏 on 2017/7/5.
//  Copyright (c) 2017 甘文鹏. All rights reserved.
//

package request

import (
	"errors"
)

// Top API: alibaba.aliqin.fc.sms.num.send
// 详细参数请参考官方Api文档：https://api.alidayu.com/docs/api.htm?apiId=25450
type AlibabaAliqinFcSmsNumSendRequest struct {
	/**
	 * 公共回传参数，在“消息返回”中会透传回该参数；举例：用户可以传入自己下级的会员ID，在消息返回时，该会员ID会包含在内，用户可以根据该会员ID识别是哪位会员使用了你的应用
	 **/
	Extend string `json:"extend"`

	/**
	 * 短信接收号码。支持单个或多个手机号码，传入号码为11位手机号码，不能加0或+86。群发短信需传入多个号码，以英文逗号分隔，一次调用最多传入200个号码。示例：18600000000,13911111111,13322222222
	 **/
	RecNum string `json:"rec_num"`

	/**
	 * 短信签名，传入的短信签名必须是在阿里大鱼“管理中心-短信签名管理”中的可用签名。如“阿里大鱼”已在短信签名管理中通过审核，则可传入”阿里大鱼“（传参时去掉引号）作为短信签名。短信效果示例：【阿里大鱼】欢迎使用阿里大鱼服务。
	 **/
	SmsFreeSignName string `json:"sms_free_sign_name"`

	/**
	 * 短信模板变量，传参规则{"key":"value"}，key的名字须和申请模板中的变量名一致，多个变量之间以逗号隔开。示例：针对模板“验证码${code}，您正在进行${product}身份验证，打死不要告诉别人哦！”，传参时需传入{"code":"1234","product":"alidayu"}
	 **/
	SmsParam string `json:"sms_param"`

	/**
	 * 短信模板ID，传入的模板必须是在阿里大鱼“管理中心-短信模板管理”中的可用模板。示例：SMS_585014
	 **/
	SmsTemplateCode string `json:"sms_template_code"`

	/**
	 * 短信类型，传入值请填写normal
	 **/
	SmsType string `json:"sms_type"`
}

func NewAlibabaAliqinFcSmsNumSendRequest() *AlibabaAliqinFcSmsNumSendRequest {
	req := new(AlibabaAliqinFcSmsNumSendRequest)
	req.SmsType = "normal"
	return req
}

func (req *AlibabaAliqinFcSmsNumSendRequest) GetMethodName() string {
	return "alibaba.aliqin.fc.sms.num.send"
}

func (req *AlibabaAliqinFcSmsNumSendRequest) ParamsIsValid() error {
	if len(req.SmsType) == 0 {
		return errors.New("sms_type is required")
	}

	if len(req.SmsFreeSignName) == 0 {
		return errors.New("sms_free_sign_name is required")
	}

	if len(req.RecNum) == 0 {
		return errors.New("rec_num is required")
	}

	if len(req.SmsTemplateCode) == 0 {
		return errors.New("sms_template_code is required")
	}

	return nil
}
