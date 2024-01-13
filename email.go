package main

import (
	"errors"
	"fmt"
	"net/smtp"
	"os"
)

func authorize() (auth smtp.Auth, sender *string, err error) {

	pass, ok := os.LookupEnv("EMAIL_PASSWORD") // this gets the password from the local computer's environment. N.B. this has to be an app password for google
	if !ok {
		return nil, nil, errors.New("no google password set. try `export EMAIL_PASSWORD=theapppassword`")
	}
	sndr, ok := os.LookupEnv("EMAIL_SENDER")
	if !ok {
		return nil, nil, errors.New("no sender email set. try `export EMAIL_SENDER=sendersemail`")
	}
	return smtp.PlainAuth("", sndr, pass, "smtp.gmail.com"), &sndr, nil
}

func email(recipient, giver person, auth smtp.Auth, sender string) error {

	to := []string{giver.Email}
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: Christmas assignment\r\nYou are getting a gift for %s.", giver.Email, recipient.Name))

	err := smtp.SendMail("smtp.gmail.com:587", auth, sender, to, msg)

	return err
}
