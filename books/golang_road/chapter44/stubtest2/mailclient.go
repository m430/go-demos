package faketest1

import (
	"go-demos/books/golang_road/chapter44/stubtest2/mailer"
	"go-demos/books/golang_road/chapter44/stubtest2/sign"
)

type mailClient struct {
	mlr mailer.Mailer
}

func New(mlr mailer.Mailer) *mailClient {
	return &mailClient{
		mlr: mlr,
	}
}

var getSign = sign.Get

func (c mailClient) ComposeAndSend(subject string, sender string, destinations []string, body string) (string, error) {
	signText := getSign(sender)
	newBody := body + "\n" + signText

	for _, dest := range destinations {
		err := c.mlr.SendMail(subject, sender, dest, newBody)
		if err != nil {
			return "", err
		}
	}
	return newBody, nil
}
