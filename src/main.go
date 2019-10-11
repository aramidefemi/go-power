package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/aramidefemi/go-power/src/helpers"
)

func sayHi (c *gin.Context) { 
	helpers.SendSms(c)
}
func main ()  {
	fmt.Printf("Hello Visitor!!")

	app := gin.Default();

	app.GET("/", sayHi)
	app.Run();
}