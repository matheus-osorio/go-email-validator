package sender

import (
	"os"

	gomail "gopkg.in/gomail.v2"
)

func SendVerification(destiny string) {

	msg := gomail.NewMessage()
	msg.SetHeader("From", os.Getenv("EMAIL_ADDR"))
	msg.SetHeader("To", destiny)
	msg.SetHeader("Subject", "Please Verify your email")
	msg.SetBody("text/html", "You would be able to check this, if this wasn't a test")

	dialer := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_ADDR"), os.Getenv("EMAIL_PASSWORD"))

	// Send the email
	if err := dialer.DialAndSend(msg); err != nil {
		panic(err)
	}

}
