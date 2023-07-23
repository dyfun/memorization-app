package Helper

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"path/filepath"
)

func SendEmail(to []string, subject string, templatePath string, data interface{}) {
	auth := smtp.PlainAuth("", os.Getenv("MAIL_FROM"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST"))

	baseDir, _ := os.Getwd()
	baseTemplatePath := filepath.Join(baseDir, templatePath)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	htmlTemplate := ParseTemplate(baseTemplatePath, data)
	title := "Subject:" + subject

	msg := []byte(title +
		"\r\n" +
		mime +
		htmlTemplate)

	err := smtp.SendMail(os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"), auth, os.Getenv("MAIL_FROM"), to, msg)

	if err != nil {
		log.Fatalln("Error:", err.Error())
	}

}

func ParseTemplate(templateFileName string, data interface{}) string {
	t, err := template.ParseFiles(templateFileName)

	if err != nil {
		log.Fatalln(err.Error())
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		log.Fatalln("Error:", err.Error())
	}

	return buf.String()
}
