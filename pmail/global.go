package pmail

var mail_smtp string
var mail_user string
var mail_pass string
var mail_port int

/*
	查看源码/usr/local/go/src/net/mail/message.go:Address结构知道
	from格式如下:
		An address such as "Barry Gibbs <bg@example.com>" is represented
		RFC 5322 address 不支持中文, 仅支持 RFC 5322格式
	示例:
		var mail_from string = "Alex Name <alex@exapmle.cn>" //不支持中文
*/
var mail_from string //不支持中文

func InitMailConfig(smtp, user, pass, from string, port int) {
	mail_smtp = smtp
	mail_user = user
	mail_pass = pass
	mail_from = from
	mail_port = port
}
