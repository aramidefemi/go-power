package controllers

import (
	"github.com/aramidefemi/go-power/src/helpers"
)


// SendSmsTo sends sms `to` number  
func SendSmsTo(to string) { 
	var advice string = helpers.FindSms()
	helpers.SendSms(advice,to)
}
