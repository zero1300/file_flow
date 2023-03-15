package email

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
)

func SendVerificationCode(to string, code string) error {
	from := viper.GetString("email.from")
	//addr := viper.GetString("email.addr")
	//username := viper.GetString("email.username")
	//password := viper.GetString("email.password")
	//host := viper.GetString("email.host")
	e := email.NewEmail()
	e.From = "File Flow  " + from
	e.To = []string{to}
	e.Subject = "File Flow 验证码"
	e.HTML = []byte(fmt.Sprintf("<h1>您的验证码: %s</h1>", code))
	//err := e.SendWithTLS(addr, smtp.PlainAuth("", username, password, host),
	//	&tls.Config{InsecureSkipVerify: true, ServerName: host})

	return nil
}
