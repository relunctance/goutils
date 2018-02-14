# pmail使用示例:

```go
	pmail.InitMailConfig("smtp@xxx.com", "username", "passwd", "fromname <from@xxx.com>", "port")
	pm := pmail.NewPmail()
	pm.AddCcUsers(map[string]string{"xxxx@abc.cn": "cc人员"})
	address := []string{"address@xxx.cn"}                                         //实际发送地址
	body := "<span style='color:green;font-size:12px;'>Hello this is test</span>" //支持html
	title := "测试Pmail邮件"                                                          //标题
	err := pm.SendMail(address, title, body)
	if err != nil {
		panic(err)
	}
```
