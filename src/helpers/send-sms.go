package helpers

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aramidefemi/go-power/src/config"
	"github.com/gin-gonic/gin"
	"github.com/prprprus/scheduler"
)

var URL string = "https://api.twilio.com/2010-04-01/Accounts/AC407b314f5a77b86f3605c3fa46fecb72/Messages.json"

func keepAppAlive() {
	println("I am runnning task.")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://go-power.herokuapp.com/", nil)
	resp, err := client.Do(req)
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	println("body", resp.Status, s, err)
}
// entry point of the app
func StartApp(c *gin.Context) {
	dt := time.Now()
	s, err := scheduler.NewScheduler(1000)
	if err != nil {
		println(err) // just example
	}
	s.Delay().Minute(1).Do(keepAppAlive)

    var message string
    message = "Sent from golang using scheduler time is: "
	s.Delay().Minute(1).Do(sendSms, message+dt.String())
	c.JSONP(200, "App Has Started Yeeeee")
}
func sendSms(message string) {
	client := &http.Client{}
	data := url.Values{}
	data.Set("Body", message)
	data.Set("From", "+12563636274")
	data.Set("To", "+2348105460858")
	req, err := http.NewRequest("POST", URL, strings.NewReader(data.Encode()))
	req.SetBasicAuth(config.AccountSid, config.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	bodyText, err := ioutil.ReadAll(resp.Body)

	s := string(bodyText)
	println("body", resp.Status, s, err)
}
