# alidayu-go

阿里大鱼短信服务 Golang SDK，代码结构简单，可扩展其他接口

## 介绍
阿里大鱼在国内算是比较好的短信服务平台了，作为阿里出品自然为开发者们提供了许多语言的SDK，但其中并不包括Golang，这也是alidayu-go产生的原因。
在使用alidayu-go的时候不能脱离官方文档，因为我们经常需要查询字段信息，传送门 => [阿里大鱼](https://api.alidayu.com/docs/api.htm?apiId=25450)

## 安装
```
go get github.com/gwpp/alidayu-go
```

## 使用
该SDK目前只支持【短信发送】、【短信发送记录查询】、【语音通知】，但代码接口清晰，可自行扩展其他需要的API，扩展说明稍后会提到

- 短信发送
    ```
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

    ```

- 短信发送查询
    ```
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

    ```

- 语音通知
    ```
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

    ```

***具体示例代码请参考[alidayu_test.go](alidayu_test.go)***

## 扩展
这个SDK的关键点不是实现了阿里大鱼的3个API，而是给使用者提供了一个可以自由扩展的框架，使用者只需要按照一套固定的格式写几十行代码就能够实现其他的API，扩展方法请看这篇博客[阿里大鱼Golang SDK —— alidayu-go](http://www.jianshu.com/p/8c5c6fcc82d4)