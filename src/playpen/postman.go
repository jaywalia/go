// send emails

package main

import (
	"fmt"
	"log"
	"net/smtp"
)

const (
	smtpHost = "smtp.gmail.com"
	smtpPort = "587"
)

func main() {
	from := "jatinderasingh@gmail.com"

	to := []string{
		"jatindera_walia@hotmail.com",
	}

	message := []byte("Test email using golang!")

	auth := smtp.PlainAuth("", from, "", smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully!")
}
