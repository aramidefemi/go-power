package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func sayHi (c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World! say hi to go!"})
}
func main ()  {
	fmt.Printf("Hello Visitor!!")

	app := gin.Default();

	app.GET("/", sayHi)
	app.Run();
}