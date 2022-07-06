package mail_test

import (
	mail "go-demos/books/golang_road/chapter30/send_mail_with_disclaimer/v1"
	"net/smtp"
	"testing"
)

func TestSendMail(t *testing.T) {
	err := mail.SendMailWithDisclaimer("gopher mail test v1", "minzheng1988@163.com", []string{"360475097@qq.com"}, "hello, gopher", "smtp.163.com:25", smtp.PlainAuth("", "minzheng1988@163.com", "zz7366231", "smtp.163.com"))
	if err != nil {
		t.Fatalf("want: nil, actual : %s\n", err)
	}
}
