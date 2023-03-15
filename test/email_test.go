package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Jordan Wright <mefqaq419093@163.com>"
	e.To = []string{"2806374351@qq.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("<h1>您的验证码: 0451 目标邮箱是我的测试账号,仅用于开发测试,并没骚扰,请知悉。</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "mefqaq419093@163.com", "DUDPJZGWYWEOMHWT", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
