package routes

import (
	controllers "github.com/kapbyte/golang-jwt-project/controllers"
	"github.com/kapbyte/golang-jwt-project/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	// incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
