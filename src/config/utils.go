package config

import ( 
	"net/http" 
)

// TwilloURL basic sms url
var TwilloURL string = "https://api.twilio.com/2010-04-01/Accounts/AC407b314f5a77b86f3605c3fa46fecb72/Messages.json"

// Client for http requests
var Client *http.Client = &http.Client{}

// AccountSid for twillo
var AccountSid string ; 

// AuthToken for twillo
var AuthToken string ;

// TwilloNumber used in twillo smses
var TwilloNumber string ;
