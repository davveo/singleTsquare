package email

import (
	"gopkg.in/gomail.v2"
)

var (
	USERNAME = "zhangqiang@xxxx.com"
	PASSWORD = "zhangqiang@xxxx.com"
	HOST     = "smtp.mxhichina.com"
	PORT     = 3306
)

func Send(mailTo, subject, body string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", USERNAME)
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(HOST, PORT, USERNAME, PASSWORD)
	return d.DialAndSend(m)
}
