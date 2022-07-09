package faketest1

import (
	"go-demos/books/golang_road/chapter44/faketest1/mailer"
	"go-demos/books/golang_road/chapter44/faketest1/sign"
)

type mailClient struct {
	mlr mailer.Mailer
}

func New(mlr mailer.Mailer) *mailClient {
	return &mailClient{
		mlr: mlr,
	}
}

func (c mailClient) ComposeAndSend(subject string, destinations []string, body string) (string, error) {
	signText := sign.Get()
	newBody := body + "\n" + signText

	for _, dest := range destinations {
		err := c.mlr.SendMail(subject, dest, newBody)
		if err != nil {
			return "", err
		}
	}
	return newBody, nil
}
