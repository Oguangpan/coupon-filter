package sendmail

import (
	"gopkg.in/gomail.v2"
)

func Send(from string, pws string, to string, dv string) (err error) {

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "最新打折优惠卷推荐")
	m.SetBody("text/html", dv)

	d := gomail.NewDialer("smtp.qq.com", 465, from, pws)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
		return err
	}
	return nil
}
