package helpers

import (
    "github.com/sfreiberg/gotwilio"
    "github.com/aramidefemi/go-power/src/config/twillo"
)
 
func SendSms() {

	twilio := gotwilio.NewTwilioClient(AccountSid, AuthToken)

	from := "+2348105460858"
	to := "+2348105460858"
	message := "Welcome to gotwilio!"
	twilio.SendSMS(from, to, message, "", "")
}