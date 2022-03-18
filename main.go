package main

import (
	"os"

	"github.com/ericklima-ca/mailmango/broker"
	"github.com/ericklima-ca/mailmango/mailer"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	var (
		EMAIL_ADDR   = os.Getenv("EMAIL_ADDR")
		EMAIL_PASS   = os.Getenv("EMAIL_PASS")
		RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
	)
	
	ms := mailer.MailerService{
		HostPort: mailer.HOSTDEFAULT,
		User:     EMAIL_ADDR,
		Passcode: EMAIL_PASS,
	}

	var worker = broker.WorkerMQ{
		Mailer:    &ms,
		HostPort:  RABBITMQ_URL,
		QueueName: "mail",
	}
	worker.StartConsume()
}
