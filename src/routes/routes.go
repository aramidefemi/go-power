package routes

import ( 
	"github.com/gin-gonic/gin" 
	"github.com/aramidefemi/go-power/src/controllers"
)

// AppRoutes declares apps routes
func AppRoutes (app *gin.Engine)  {  
	app.GET("/", controllers.GreetApp)
	app.GET("/motun", func (c *gin.Context) { controllers.SendSmsTo("+2348100026809") } ) 
}
