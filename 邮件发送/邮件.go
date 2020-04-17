package main

import (
	"fmt"
	"net/smtp"
)

type client struct {
	User     string
	Password string
	Addr     string
	Port     int
}

type email struct {
	Sender   string
	Receiver []string
	Subject  string
	Body     string
	Type     string
}

func (cli client) Send(email email) error {

	auth := smtp.PlainAuth("", cli.User, cli.Password, cli.Addr)

	var contentType string
	switch email.Type {
	case "html":
		contentType = fmt.Sprintf("text/%s; charset=UTF-8", email.Type)
	case "plain":
		contentType = fmt.Sprintf("text/%s; charset=UTF-8", email.Type)
	default:
		return fmt.Errorf(fmt.Sprintf("the type(%s) of the email can't be distinguished", email.Type))
	}

	// msg只是用来控制显示的,
	// To字段: 显示要发送给谁(可以不要)
	// From字段: 显示由谁发(可以不要)
	// Subject字段: 控制标题(可以不要)
	// Content-Type: 指定Body内容的类型，默认为plain。
	// Body: 邮件的内容
	msg := []byte(fmt.Sprintf("From: %s\r\nSubject: %s\r\nContent-Type: %s\r\n\r\n%s", email.Sender, email.Subject, contentType, email.Body))

	host := fmt.Sprintf("%s:%d", cli.Addr, cli.Port)
	return smtp.SendMail(host, auth, cli.User, email.Receiver, msg)
}

func main() {
	client := client{
		"417165709@qq.com",
		"zdvvnrxchbfgbhja",
		"smtp.qq.com",
		25,
	}
	email := email{
		"417165709@qq.com",
		[]string{"417165709@qq.com","2273454787@qq.com"},
		"使用Golang发送邮件",
		`
      <html>
      	<body>
      		<h3>"Test send to email"</h3>
      	</body>
      </html>
      `,
		"html",
	}

	if err := client.Send(email); err != nil {
		fmt.Println(err)
	}
}
