package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/aramidefemi/go-power/src/helpers/send-email"
)

func sayHi (c *gin.Context) {
	SendEmail()
	c.JSON(200, gin.H{
		"message": "Hello World! say hi to go!"})
}
func main ()  {
	fmt.Printf("Hello Visitor!!")

	app := gin.Default();

	app.GET("/", sayHi)
	app.Run();
}