package helpers

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/aramidefemi/go-power/src/config"
	"github.com/gin-gonic/gin"
	"github.com/sfreiberg/gotwilio"
)

var URL string = "https://api.twilio.com/2010-04-01/Accounts/AC407b314f5a77b86f3605c3fa46fecb72/Messages.json"

func SendSms(c *gin.Context) {
	var params map[string][]string

	params = c.Request.URL.Query()

	println("got here with", params["message"])
	data := url.Values{}
	data.Set("Body", "seent from golang")
	data.Set("From", "+12563636274")
	data.Set("To", "+2348105460858")
	twilio := gotwilio.NewTwilioClient(config.AccountSid, config.AuthToken)
	from := "+12563636274"
	to := "+2348105460858"
	message := "Welcome to gotwilio!"
	twilio.SendSMS(from, to, message, "sent!!!!", "")

	client := &http.Client{}
	//pass the values to the request's body
	req, err := http.NewRequest("POST", URL, strings.NewReader(data.Encode()))
	req.SetBasicAuth(config.AccountSid, config.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	bodyText, err := ioutil.ReadAll(resp.Body)

	s := string(bodyText)
	println("body", resp.Status, s, err)

	c.JSONP(200, s)
}
