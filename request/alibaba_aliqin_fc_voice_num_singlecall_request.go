//
//
//  alidayu-go
//
//  Created by 甘文鹏 on 2017/7/5.
//  Copyright (c) 2017 甘文鹏. All rights reserved.
//

package request

import "errors"

// Top API: alibaba.aliqin.fc.voice.num.singlecall
// 详细参数请参考官方Api文档：https://api.alidayu.com/docs/api.htm?apiId=25445
type AlibabaAliqinFcVoiceNumSinglecallRequest struct {
	Extend        string `json:"extend"`
	CalledNum     string `json:"called_num"`
	CalledShowNum string `json:"called_show_num"`
	VoiceCode     string `json:"voice_code"`
}

func NewAlibabaAliqinFcVoiceNumSinglecallRequest() *AlibabaAliqinFcVoiceNumSinglecallRequest {
	return new(AlibabaAliqinFcVoiceNumSinglecallRequest)
}

func (req *AlibabaAliqinFcVoiceNumSinglecallRequest) GetMethodName() string {
	return "alibaba.aliqin.fc.voice.num.singlecall"
}

func (req *AlibabaAliqinFcVoiceNumSinglecallRequest) ParamsIsValid() error {
	if len(req.CalledNum) == 0 {
		return errors.New("called_num is required")
	}
	if len(req.CalledShowNum) == 0 {
		return errors.New("called_show_num is required")
	}
	if len(req.VoiceCode) == 0 {
		return errors.New("voice_code is required")
	}

	return nil
}
