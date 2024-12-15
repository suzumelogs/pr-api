package routes

import (
	"pr-api/controllers/admin"
	"pr-api/controllers/auth"
	"pr-api/controllers/client"
	"pr-api/controllers/moderator"
	"pr-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/signup", auth.Signup)
	}

	protectedGroup := r.Group("/")
	protectedGroup.Use(middlewares.LoginRequired)
	{
		adminGroup := protectedGroup.Group("/admin")
		adminGroup.Use(middlewares.AdminRoleRequired)
		{
			adminGroup.GET("/dashboard", admin.AdminController)
		}

		moderatorGroup := protectedGroup.Group("/moderator")
		moderatorGroup.Use(middlewares.ModeratorRoleRequired)
		{
			moderatorGroup.GET("/dashboard", moderator.ModeratorController)
		}

		clientGroup := protectedGroup.Group("/client")
		clientGroup.Use(middlewares.ClientRoleRequired)
		{
			clientGroup.GET("/dashboard", client.ClientController)
		}
	}
}