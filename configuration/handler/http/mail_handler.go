package http

import (
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Voucher Labs <voucherlabs.official@gmail.com>"
const CONFIG_AUTH_EMAIL = "voucherlabs.official@gmail.com"
const CONFIG_AUTH_PASSWORD = "gpIvwI7fYbh9Bd1D5g+gnmu/hcT4YkgPzt2Inpjrjuw="

func GenerateWIB(t time.Time) time.Time {
	wib, err := time.LoadLocation("Asia/Jakarta")
	if err == nil {
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), wib)
	}
	return t
}
func Sendmail(form CreateCredentialBodyRequest) {
	var text string
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", form.Email, "voucherlabs.official@gmail.com")
	mailer.SetAddressHeader("Cc", "voucherlabs.official@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Email Verification")
	if form.Seller {
		text = "You're currently a seller"
	} else {
		text = "You're currently a buyer"
	}
	mailer.SetBody("text/html", "Hello, "+form.Username+"<br>have a nice day, here's your credentials : <br> Email : "+form.Email+"<br> Password : "+form.Password+"<br><br>"+text)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Print(dialer)
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}

func verify_mail(form VerifyResult) {
	var link, text string
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", form.Email, "voucherlabs.official@gmail.com")
	mailer.SetAddressHeader("Cc", "voucherlabs.official@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Code Verification")
	link = os.Getenv("RUNNING_HOST") + form.Credential_id.String()
	text = "Please, click this link before " + form.Expiresat.Format(time.RFC1123)

	mailer.SetBody("text/html", "Here's your verification link, click link<br><a href='"+link+"'>Click here</a><br>"+text)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Print(dialer)
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}

func forgot_mail(id uuid.UUID, email string) {
	var link, text string
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email, "voucherlabs.official@gmail.com")
	mailer.SetAddressHeader("Cc", "voucherlabs.official@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Code Verification")
	link = os.Getenv("RUNNING_HOST") + "reset-password/" + id.String()

	mailer.SetBody("text/html", "Please click this link to reset your password<br><a href='"+link+"'>Click here</a><br>"+text)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Print(dialer)
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
