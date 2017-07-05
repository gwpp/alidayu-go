//
//
//  alidayu-go
//
//  Created by 甘文鹏 on 2017/7/5.
//  Copyright (c) 2017 甘文鹏. All rights reserved.
//

package request

import "errors"

// Top API: alibaba.aliqin.fc.sms.num.query
// 详细参数请参考官方Api文档：https://api.alidayu.com/docs/api.htm?apiId=26039
type AlibabaAliqinFcSmsNumQueryRequest struct {
	BizId       string `json:"biz_id"`
	RecNum      string `json:"rec_num"`
	QueryDate   string `json:"query_date"`
	CurrentPage string `json:"current_page"`
	PageSize    string `json:"page_size"`
}

func NewAlibabaAliqinFcSmsNumQueryRequest() *AlibabaAliqinFcSmsNumQueryRequest {
	return new(AlibabaAliqinFcSmsNumQueryRequest)
}

func (req *AlibabaAliqinFcSmsNumQueryRequest) GetMethodName() string {
	return "alibaba.aliqin.fc.sms.num.query"
}

func (req *AlibabaAliqinFcSmsNumQueryRequest) ParamsIsValid() error {
	if len(req.RecNum) == 0 {
		return errors.New("rec_num is required")
	}
	if len(req.QueryDate) == 0 {
		return errors.New("query_date is required")
	}
	if len(req.CurrentPage) == 0 {
		return errors.New("current_page is required")
	}
	if len(req.PageSize) == 0 {
		return errors.New("page_size is required")
	}

	return nil
}
