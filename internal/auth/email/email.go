package email1

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/unknownn17/Internship_Task/internal/config"
	"gopkg.in/gomail.v2"
)

func sendEmail(to, subject, body string) error {
	c:=config.Configuration()
	m := gomail.NewMessage()
	m.SetHeader("From", c.Email.Sender)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, c.Email.Sender, c.Email.Password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func Sent(email string) (string,error) {
	code, err := generateRandomCode()
	if err != nil {
		log.Printf("Failed to generate code: %v", err)
		return "",err
	}
	to := email
	subject := "Your Confirmation Code"
	body := fmt.Sprintf("Your confirmation code is: %s", code)

	if err := sendEmail(to, subject, body); err != nil {
		log.Printf("Failed to send email: %v", err)
		return "",err
	}

	log.Println("Email sent successfully")
	return code,nil
}

func generateRandomCode() (string, error) {
	max := big.NewInt(1000000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	code := fmt.Sprintf("%06d", n.Int64())
	return code, nil
}
