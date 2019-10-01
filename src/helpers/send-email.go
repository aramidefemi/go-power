package helpers

import (
	"github.com/kataras/go-mailer"
	"github.com/aramidefemi/go-power/src/config/mailgun"
)

	// SendEmail sends emails out
func SendEmail() {
    // sender configuration.
    config := MailConfig

    // initalize a new mail sender service.
    sender := mailer.New(config)

    // the subject/title of the e-mail.
    subject := "Hello subject"

    // the rich message body.
    content := `<h1>Hello</h1> <br/><br/> <span style="color:red"> This is the rich message body </span>`

    // the recipient(s).
    to := []string{"kataras2006@hotmail.com", "kataras2018@hotmail.com"}

    // send the e-mail.
    err := sender.Send(subject, content, to...)

    if err != nil {
        println("error while sending the e-mail: " + err.Error())
    }
}