package mailer

type Mailer interface {
	SendMail(subject, sender, destination, body string) error
}
