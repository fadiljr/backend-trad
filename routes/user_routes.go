// routes/user_routes.go
package routes

import (
	"go-gorm-postgresql/controllers"
	"go-gorm-postgresql/middleware"

	"github.com/gin-gonic/gin"
)

// UserRoutes menangani routing untuk user
func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	userGroup.Use(middleware.JWTAuthMiddleware())

	userGroup.GET("/", controllers.GetAllUsers)
	userGroup.GET("/:id", controllers.GetUser)
	userGroup.POST("/", controllers.CreateUser)
	userGroup.PUT("/:id", controllers.UpdateUser)
	userGroup.DELETE("/:id", controllers.DeleteUser)

	router.GET("/profile", middleware.JWTAuthMiddleware(), controllers.GetProfile)
}
