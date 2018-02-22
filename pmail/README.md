# pmail使用示例:

```go
    // smtp = "smtp@xxx.com"
	pmail.InitMailConfig("smtp@xxx.com", "username", "passwd", "fromname <from@xxx.com>", "port")
	pm := pmail.NewPmail()
	pm.AddCcUsers(map[string]string{"xxxx@abc.cn": "cc人员"})   //[可选] 抄送人员
	address := []string{"address@xxx.cn"}                                         //实际发送地址
	body := "<span style='color:green;font-size:12px;'>Hello this is test</span>" //支持html
	title := "测试Pmail邮件"                                                          //标题
    pm.AddAttachment("/tmp/test_mail.txt"); //[可选] 增加附件
	err := pm.SendMail(address, title, body)     //内部实现支持HTML
	if err != nil {
		panic(err)
	}
```
