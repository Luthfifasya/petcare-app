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

	// Pet
	pet := api.Group("/pets")
	pet.Use(middleware.JWTAuth())
	{
		pet.GET("", controllers.GetPets)
		pet.POST("", controllers.CreatePet)
		pet.GET("/:id", controllers.GetPetByID)
		pet.PUT("/:id", controllers.UpdatePet)
		pet.DELETE("/:id", controllers.DeletePet)
	}

	appointment := api.Group("/appointment")
	appointment.Use(middleware.JWTAuth())
	{
		appointment.GET("", controllers.GetAppointments)
		appointment.POST("", controllers.CreateAppointment)
		appointment.GET("/:id", controllers.GetAppointmentByID)
		appointment.PUT("/:id", controllers.UpdateAppointment)
		appointment.DELETE("/:id", controllers.DeleteAppointment)
	}
}
