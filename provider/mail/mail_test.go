package mail

import (
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"github.com/purwantogz/go-impact/config"
)

func TestSend(t *testing.T) {
	err := godotenv.Load()
	mailPort, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	config.MailConf = &config.MailConfig{
		Driver:   os.Getenv("MAIL_DRIVER"),
		Host:     os.Getenv("MAIL_HOST"),
		Port:     mailPort,
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
	}

	mailServe := Build("purwanto.dev@gmail.com", "test_go_email", "asdhs shda shds dhdas")

	err = mailServe.Send()

	if err != nil {
		t.Errorf(err.Error())
	}
}
