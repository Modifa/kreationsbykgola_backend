package services

import (
	"Modifa/kreationsbykgola_backend/models"
	utils "Modifa/kreationsbykgola_backend/utils"
	v "Modifa/kreationsbykgola_backend/views"
	"strings"

	"gopkg.in/gomail.v2"
)

/*Send ... */
func Send(Order models.Order) {

	from := utils.MustGetenv("EmailFrom")

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", Order.Email)
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Proect Main")
	m.SetBody("text/html", Order.Email)
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "projectcommunication123@gmail.com", "Modifa1993*")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func OrderRecieved(To string, Message string, FullName string) {

	from := utils.MustGetenv("EmailFrom")
	vr := *new(v.Messages).New()

	subject := strings.ReplaceAll(vr.MessageSubject, "{FullName}", FullName)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", To)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", Message)
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("127.0.0.1", 1143, "kreationsbykgola@protonmail.com", "Modifa1993*")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func OrderRecievedResponse(To string) {

	from := utils.MustGetenv("EmailFrom")
	vr := *new(v.Messages).New()

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", To)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", vr.NoneReplyResponse)
	m.SetBody("text/html", vr.OrderResponse)
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "projectcommunication123@gmail.com", "Modifa1993*")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
