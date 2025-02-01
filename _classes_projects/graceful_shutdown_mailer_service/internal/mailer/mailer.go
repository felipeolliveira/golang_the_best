package mailer

import (
	"log/slog"
	"sync"
	"time"

	"github.com/go-mail/mail/v2"
)

type Sender struct {
	dialer *mail.Dialer
	addr   string
}

func NewSender(host string, port int, username, password, addr string) Sender {
	return Sender{
		dialer: mail.NewDialer(host, port, username, password),
		addr:   addr,
	}
}

func (s Sender) Send(recipient string, wg *sync.WaitGroup) error {
	defer wg.Done()

	msg := mail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", s.addr)
	msg.SetHeader("Subject", "An Email from a Xablau SMTP wrapper")
	msg.SetHeader("text/plain", "Hey this as test email from golang mailer. I hope you get this... xablau")

	slog.Info("mailer", "Started sending", recipient)
	defer slog.Info("mailer", "Finished sending", recipient)

	time.Sleep(5 * time.Second)

	return s.dialer.DialAndSend(msg)
}
