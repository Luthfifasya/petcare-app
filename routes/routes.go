package routes

import (
	"petcare-app/controllers"
	"petcare-app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.POST("/users/register", controllers.RegisterUser)
	api.POST("/users/login", controllers.LoginUser)
	api.GET("/users", middleware.JWTAuth(), controllers.GetUsers)
}
