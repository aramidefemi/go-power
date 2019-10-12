package main

import ( 
	"github.com/gin-gonic/gin"
	"github.com/aramidefemi/go-power/src/controllers"
	"github.com/aramidefemi/go-power/src/routes"
)
 

func main ()  { 
	app := gin.Default();
	routes.AppRoutes(app)
	controllers.StartApp()
	app.Run();
}
