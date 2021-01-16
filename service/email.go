package service

import (
	"strconv"

	"github.com/unknwon/goconfig"
	"gopkg.in/gomail.v2"
)

type Email struct {
	host string
	port string
	user string
	pass string
}

func NewEmail(conf *goconfig.ConfigFile) *Email {
	// host := conf.MustValue("smtp", "email_host", "")
	// port := conf.MustValue("smtp", "port", "")
	// user := conf.MustValue("smtp", "email_user", "")
	// pass := conf.MustValue("smtp", "email_pwd", "")
	// host := "smtp.gmail.com"
	// port := "587"
	// user := "applestore0851@gmail.com"
	// pass := "dzanhlboleydzraq"
	host := "smtp.qq.com"
	port := "465"
	user := "qsyxcx@qq.com"
	pass := "gkdzjdwerowvhafe"

	return &Email{host: host, port: port, user: user, pass: pass}
}

func (this *Email) Send(mailTo []string, subject string, body string) error {
	port, _ := strconv.Atoi(this.port)
	m := gomail.NewMessage()
	// 发件人信息
	m.SetHeader("From", "<"+this.user+">")
	// 收件人
	m.SetHeader("To", mailTo...)
	// 主题
	m.SetHeader("Subject", subject)
	// 内容
	m.SetBody("text/html", body)
	// 构造附件信息，同时对附件进行重命名
	name := "ck.txt"
	m.Attach("cookie.txt",
		gomail.Rename(name),
	)

	d := gomail.NewDialer(this.host, port, this.user, this.pass)
	// log.Warn("正在发送通知...")
	err := d.DialAndSend(m)
	if err != nil {
		//log.Error("邮件发送失败，返回错误:" + err.Error())
	} else {
		// log.Println("邮件发送成功")
	}
	return err
}
