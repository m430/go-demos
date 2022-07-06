package mail_test

import (
	mail "go-demos/books/golang_road/chapter30/send_mail_with_disclaimer/v2"
	"net/smtp"
	"testing"
)

type FakeEmailServer struct {
	subject string
	from    string
	to      []string
	content string
}

func (s *FakeEmailServer) Send(subject, from string, to []string, content string, mailserver string, a smtp.Auth) error {
	s.subject = subject
	s.from = from
	s.to = to
	s.content = content
	return nil
}

func TestSendMail(t *testing.T) {
	s := &FakeEmailServer{}
	err := mail.SendMailWithDisclaimer(s, "gopher mail test v2", "minzheng1988@163.com", []string{"360475097@qq.com"}, "hello, gopher", "smtp.163.com:25", smtp.PlainAuth("", "minzheng1988@163.com", "zz7366231", "smtp.163.com"))
	if err != nil {
		t.Fatalf("want: nil, actual : %s\n", err)
	}
	want := "hello, gopher" + "\n\n" + mail.DISCLAIMER
	if s.content != want {
		t.Fatalf("want: %s, actual: %s\n", want, s.content)
	}
}
