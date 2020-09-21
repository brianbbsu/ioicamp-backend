package main

import (
	"bytes"
	"html/template"
	"path"

	gomail "gopkg.in/gomail.v2"
)

func sendMail(mail *gomail.Message) error {
	mail.SetHeader("From", Config.GetString("email.email"))
	d := gomail.NewDialer(Config.GetString("email.server"), Config.GetInt("email.port"),
		Config.GetString("email.email"), Config.GetString("email.password"))

	err := d.DialAndSend(mail)
	if err != nil {
		panic(err)
	}

	return nil
}

func sendTestMail(recipient string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("To", recipient)
	mail.SetHeader("Subject", "This is just a test")
	mail.SetBody("text/html", "Hello!!<br>This is a test.<br><b>How are you?</b>")

	err := sendMail(mail)
	return err
}

func sendEmailVerification(address string, token string) error {
	templatePath := path.Join(Config.GetString("TemplateDir"), "EmailVerification.html")
	emailTemplate := template.Must(template.ParseFiles(templatePath))
	data := struct {
		Email string
		Token string
	}{
		Email: address,
		Token: token,
	}
	contentBuf := new(bytes.Buffer)
	if err := emailTemplate.Execute(contentBuf, data); err != nil {
		return err
	}
	subjectBuf := new(bytes.Buffer)
	if err := emailTemplate.ExecuteTemplate(subjectBuf, "Subject", data); err != nil {
		return err
	}
	subject, content := subjectBuf.String(), contentBuf.String()

	mail := gomail.NewMessage()
	mail.SetHeader("To", address)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", content)

	err := sendMail(mail)
	return err
}
