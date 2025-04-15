package routes

import (
	"github.com/NochboolPrime/task_manager/pkg/controllers"
	"github.com/NochboolPrime/task_manager/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	public := router.Group("/api")
	{
		public.POST("/login", controllers.Login)
	}

	protected := router.Group("/api")
	protected.Use(middleware.JWTAuth())
	{
		protected.POST("/teams", controllers.CreateTeam)
		protected.POST("/projects", controllers.CreateProject)
		protected.POST("/tasks", controllers.CreateTask)
	}
}
