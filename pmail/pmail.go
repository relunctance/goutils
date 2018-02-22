package pmail

import (
	"net/smtp"
	"os"

	"github.com/go-gomail/gomail"
)

type Pmail struct {
	M *gomail.Message
	D *gomail.Dialer
}

func NewPmail() *Pmail {
	p := &Pmail{}
	p.M = gomail.NewMessage()
	p.M.SetHeader("From", mail_from) //来源主题
	return p
}

//添加抄送人员
func (p *Pmail) AddCcUsers(ccUsers map[string]string) {
	if len(ccUsers) == 0 {
		return
	}
	for mailAddress, name := range ccUsers {
		p.M.SetAddressHeader("Cc", mailAddress, name)
	}
}

//添加附件
func (p *Pmail) AddAttachment(filepath string) error {
	if _, err := os.Stat(filepath); err != nil { //文件存在的情况下, 进行添加附件
		return err
	}
	p.M.Attach(filepath)
	return nil
}

//发送邮件
func (p *Pmail) SendMail(address []string, title, body string) error {
	p.M.SetHeader("To", address...) //需要接收的人员
	p.M.SetHeader("Subject", title) //设置标题
	p.M.SetBody("text/html", body)  //发送的内容
	p.addDialer()
	if err := p.D.DialAndSend(p.M); err != nil {
		return err
	}
	return nil

}

type unencryptedAuth struct {
	smtp.Auth
}

//重新 smtp.Auth 接口中Start函数
func (a *unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)

}
func (p *Pmail) addDialer() {
	if p.D == nil {
		d := gomail.NewDialer(mail_smtp, mail_port, mail_user, mail_pass)
		d.Auth = &unencryptedAuth{ //用于修复: unencrypted connection 这个问题
			smtp.PlainAuth(
				"",
				mail_user,
				mail_pass,
				mail_smtp,
			)}
		p.D = d
	}

}
