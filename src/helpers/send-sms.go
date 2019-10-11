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
    "encoding/json"
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
	findSms()
}
// StartApp entry point of the app
func StartApp(c *gin.Context) {
    dt := time.Now()
    hour,min,sec := dt.Clock()
    
	s, err := scheduler.NewScheduler(1000)
	if err != nil {
		println(err) // just example
	}
	s.Delay().Minute(20).Do(keepAppAlive) 
    // findSms()
    println(hour,min,sec)
    var Dhour int = hour
	c.JSONP(200, "App Has Started Yeeeee: "+dt.String()+" hours: "+string(Dhour))
}
func findSms() { 
    client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.adviceslip.com/advice", nil)
	resp, err := client.Do(req)
	bodyText, err := ioutil.ReadAll(resp.Body)
    s := string(bodyText)

    var data map[string]map[string]string
    json.Unmarshal([]byte(bodyText), &data)

    println("body --- ", resp.Status, s,  err) 
    var  advice string = data["slip"]["advice"]
    sendSms(advice)
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
