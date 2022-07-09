package faketest1

import (
	"fmt"
	"testing"
)

type fakeOkMailer struct {
}

func (m *fakeOkMailer) SendMail(subject, dest, body string) error {
	return nil
}

type fakeFailMailer struct {
}

func (m fakeFailMailer) SendMail(subject, dest, body string) error {
	return fmt.Errorf("can not reach the mail server of dest [%s]", dest)
}

func TestComposeAndSendOk(t *testing.T) {
	m := &fakeOkMailer{}
	mc := New(m)
	_, err := mc.ComposeAndSend("hello, fake test", []string{"xxx@example.com"}, "the test body")
	if err != nil {
		t.Errorf("want nil, got %v", err)
	}
}

func TestComposeAndSendFail(t *testing.T) {
	m := fakeFailMailer{}
	mc := New(m)
	_, err := mc.ComposeAndSend("hello, fake test", []string{"xxx@example.com"}, "the test body")
	if err == nil {
		t.Errorf("want non-nil, got nil")
	}
}

// 总结：fake不具备在测试前对返回的对象的预置能力
