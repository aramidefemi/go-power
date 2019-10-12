package helpers

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings" 
 
	"github.com/aramidefemi/go-power/src/config" 
    "encoding/json"
)


// FindSms roams the internet for an advice message 
func FindSms() string { 
	req, err := http.NewRequest("GET", "https://api.adviceslip.com/advice", nil)
	resp, err := config.Client.Do(req)
	bodyText, err := ioutil.ReadAll(resp.Body) 
    var data map[string]map[string]string
    json.Unmarshal([]byte(bodyText), &data)
    var advice string = data["slip"]["advice"] 
	println("FindSms function: ", resp.Status, err, "result: ",advice)
	return advice;
}

// SendSms send smses using twillo api
func SendSms(message string, to string) { 
	data := url.Values{}
	data.Set("Body", message)
	data.Set("From", config.TwilloNumber)
	data.Set("To", to)
	req, err := http.NewRequest("POST", config.TwilloURL, strings.NewReader(data.Encode()))
	req.SetBasicAuth(config.AccountSid, config.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := config.Client.Do(req)
	println("SendSms function: ", resp.Status, err)
}
