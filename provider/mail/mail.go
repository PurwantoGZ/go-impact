package mail

import (
	"bytes"
	"html/template"
	"log"
	"os"

	"github.com/purwantogz/go-impact/config"
	"gopkg.in/gomail.v2"
)

type mailService struct {
	receiver string
	subject  string
	message  string
}

type info struct {
	Name string
}

//BuildWithHtml build mail service with html
func BuildWithHtml(to, subject string) IMailService {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	t, err := template.ParseFiles(wd + "/template/alert.html")
	if err != nil {
		log.Println(err)
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, &info{
		Name: "Purwanto",
	}); err != nil {
		log.Println(err)
	}
	return &mailService{
		receiver: to,
		subject:  subject,
		message:  tpl.String(),
	}
}

//Build build mail service
func Build(to, subject, message string) IMailService {
	return &mailService{
		receiver: to,
		subject:  subject,
		message:  message,
	}
}

//Send send email
func (serv *mailService) Send() error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.MailConf.Username)
	mailer.SetHeader("To", serv.receiver)
	mailer.SetHeader("Subject", serv.subject)
	mailer.SetBody("text/html", serv.message)
	//m.Attach("/home/Alex/lolcat.jpg")

	dialer := gomail.NewDialer(
		config.MailConf.Host,
		config.MailConf.Port,
		config.MailConf.Username,
		config.MailConf.Password,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	return nil
}
