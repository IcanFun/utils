package utils

import (
	"crypto/tls"
	"fmt"
	"mime"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
	"time"

	"github.com/IcanFun/utils/utils/log"
)

const (
	CONN_SECURITY_TLS      = "TLS"
	CONN_SECURITY_STARTTLS = "STARTTLS"
)

type EmailSettings struct {
	EnableSignUpWithEmail             bool
	EnableSignInWithEmail             *bool
	EnableSignInWithUsername          *bool
	FeedbackName                      string
	FeedbackEmail                     string
	FeedbackOrganization              *string
	EnableSMTPAuth                    *bool
	SMTPUsername                      string
	SMTPPassword                      string
	SMTPServer                        string
	SMTPPort                          string
	ConnectionSecurity                string
	SkipServerCertificateVerification *bool
}

func encodeRFC2047Word(s string) string {
	return mime.BEncoding.Encode("utf-8", s)
}

func connectToSMTPServer(settings EmailSettings) (net.Conn, error) {
	var conn net.Conn
	var err error

	if settings.ConnectionSecurity == CONN_SECURITY_TLS {
		tlsconfig := &tls.Config{
			InsecureSkipVerify: *settings.SkipServerCertificateVerification,
			ServerName:         settings.SMTPServer,
		}

		conn, err = tls.Dial("tcp", settings.SMTPServer+":"+settings.SMTPPort, tlsconfig)
		if err != nil {
			log.Error("connectToSMTPServer=>tls.Dial err:%s", err.Error())
			return nil, err
		}
	} else {
		conn, err = net.Dial("tcp", settings.SMTPServer+":"+settings.SMTPPort)
		if err != nil {
			log.Error("connectToSMTPServer=>net.Dial err:%s", err.Error())
			return nil, err
		}
	}

	return conn, nil
}

func newSMTPClient(conn net.Conn, settings EmailSettings) (*smtp.Client, error) {
	c, err := smtp.NewClient(conn, settings.SMTPServer+":"+settings.SMTPPort)
	if err != nil {
		log.Error(T("github.com/IcanFun/utils.mail.new_client.open.error"), err)
		return nil, err
	}

	if settings.ConnectionSecurity == CONN_SECURITY_STARTTLS {
		tlsconfig := &tls.Config{
			InsecureSkipVerify: *settings.SkipServerCertificateVerification,
			ServerName:         settings.SMTPServer,
		}
		c.StartTLS(tlsconfig)
	}

	if *settings.EnableSMTPAuth {
		auth := smtp.PlainAuth("", settings.SMTPUsername, settings.SMTPPassword, settings.SMTPServer+":"+settings.SMTPPort)

		if err = c.Auth(auth); err != nil {
			log.Error("newSMTPClient=>Auth err:%s", err.Error())
			return nil, err
		}
	}
	return c, nil
}

func SendMail(fromName, to, subject, body string, settings EmailSettings) error {
	return SendMailUsingConfig(fromName, to, subject, body, settings)
}

func SendMailUsingConfig(fromeName, to, subject, body string, settings EmailSettings) error {
	if len(settings.SMTPServer) == 0 {
		return nil
	}

	log.Debug(T("github.com/IcanFun/utils.mail.send_mail.sending.debug"), to, subject)

	fromMail := mail.Address{Name: fromeName, Address: settings.FeedbackEmail}
	toMail := mail.Address{Name: "", Address: to}

	headers := make(map[string]string)
	headers["From"] = fromMail.String()
	headers["To"] = toMail.String()
	headers["Subject"] = encodeRFC2047Word(subject)
	headers["MIME-version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"utf-8\""
	headers["Content-Transfer-Encoding"] = "8bit"
	headers["Date"] = time.Now().Format(time.RFC1123Z)

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n<html><body>" + body + "</body></html>"

	conn, err := connectToSMTPServer(settings)
	if err != nil {
		log.Error("SendMailUsingConfig=>connectToSMTPServer err:%s", err.Error())
		return err
	}
	defer conn.Close()

	c, err := newSMTPClient(conn, settings)
	if err != nil {
		log.Error("SendMailUsingConfig=>newSMTPClient err:%s", err.Error())
		return err
	}
	defer c.Quit()
	defer c.Close()

	if err := c.Mail(fromMail.Address); err != nil {
		log.Error("SendMailUsingConfig=>Mail err:%s", err.Error())
		return err
	}

	if err := c.Rcpt(toMail.Address); err != nil {
		log.Error("SendMailUsingConfig=>Rcpt err:%s", err.Error())
		return err
	}

	w, err := c.Data()
	if err != nil {
		log.Error("SendMailUsingConfig=>Data err:%s", err.Error())
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Error("SendMailUsingConfig=>Write err:%s", err.Error())
		return err
	}

	err = w.Close()
	if err != nil {
		log.Error("SendMailUsingConfig=>Close err:%s", err.Error())
		return err
	}

	return nil
}

func IsEmail(email string) bool {
	return strings.Index(email, "@") != -1
}
