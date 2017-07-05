//
//
//  alidayu-go
//
//  Created by 甘文鹏 on 2017/7/5.
//  Copyright (c) 2017 甘文鹏. All rights reserved.
//

package request

type AlibabaRequest interface {
	GetMethodName() string
	ParamsIsValid() error
}
