package controllers

import ( 
	"net/http" 
	"time"
 
	"github.com/gin-gonic/gin"
    "github.com/prprprus/scheduler" 
)

 
func keepAppAlive() {
	println("I am runnning task.")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://go-power.herokuapp.com/", nil)
	resp, err := client.Do(req) 
	println("keepAppAlive function: ", resp.Status, err)
}

func sendTextToSelf() {
	SendSmsTo("+2348105460858")
}

// StartApp entry point of the app
func StartApp() {
	dt := time.Now()
	hour,min,sec := dt.Clock() 
	s, err := scheduler.NewScheduler(1000)
	if err != nil {
		println(err)
	}
	if hour == 1 {
		sendTextToSelf()
	}
    println("App Has Started! this hour - ",hour,"this min -",min,"this sec -",sec)
	s.Delay().Minute(20).Do(keepAppAlive) 
}

// GreetApp greets our users on '/'
func GreetApp(c *gin.Context) {
	dt := time.Now()
	hour,min,sec := dt.Clock()  
    println("App Has Started! this hour - ",hour,"this min -",min,"this sec -",sec)
	c.JSONP(200, "Hello this GO! app is alive!: " + dt.String())
}
