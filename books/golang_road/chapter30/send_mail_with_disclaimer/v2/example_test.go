package mail_test

import (
	"fmt"
	mail "go-demos/books/golang_road/chapter30/send_mail_with_disclaimer/v2"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

type EmailSenderAdapter struct {
	e *email.Email
}

func (adapter *EmailSenderAdapter) Send(subject, from string, to []string, content string, mailserver string, a smtp.Auth) error {
	adapter.e.Subject = subject
	adapter.e.From = from
	adapter.e.To = to
	adapter.e.Text = []byte(content)
	return adapter.e.Send(mailserver, a)
}

func TestExampleSendMailWithDisclaimer(t *testing.T) {
	adapter := &EmailSenderAdapter{e: email.NewEmail()}
	err := mail.SendMailWithDisclaimer(adapter, "gopher mail test v2", "minzheng1988@163.com", []string{"360475097@qq.com"}, "hello, gopher", "smtp.163.com:25", smtp.PlainAuth("", "minzheng1988@163.com", "zz7366231", "smtp.163.com"))
	if err != nil {
		fmt.Printf("SendMail error: %s\n", err)
		return
	}
	fmt.Println("SendMail ok")
}
