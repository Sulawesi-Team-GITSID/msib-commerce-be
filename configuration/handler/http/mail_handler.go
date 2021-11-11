package http

import (
	"log"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Voucher Labs <voucherlabs.official@gmail.com>"
const CONFIG_AUTH_EMAIL = "voucherlabs.official@gmail.com"
const CONFIG_AUTH_PASSWORD = "gpIvwI7fYbh9Bd1D5g+gnmu/hcT4YkgPzt2Inpjrjuw="

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

func verify_mail() {

}
