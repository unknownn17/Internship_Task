package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	jwttoken "github.com/unknownn17/Internship_Task/internal/auth/jwt"
	"github.com/unknownn17/Internship_Task/internal/connections"
	_ "github.com/unknownn17/Internship_Task/internal/docs"
)

// @title           Task Management API
// @version         2.0
// @description     This is an API for managing tasks in the system.
// @host            3.127.221.197:8080
// @BasePath /
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

func Router() {
	r := gin.Default()

	handler := connections.NewHandler()

	r.POST("/user/register", handler.Register)
	r.POST("/user/verify", handler.Verify)
	r.POST("/user/login", handler.LogIn)

	protected := r.Group("/")
	protected.Use(jwttoken.JWTMiddleware())
	{
		protected.POST("/task", handler.CreateTask)
		protected.GET("/task", handler.GetTask)
		protected.GET("/tasks", handler.GetsTasks)
		protected.PUT("/task", handler.UpdateTask)
		protected.DELETE("/task", handler.DeleteTask)
	}

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(":8080")
}
