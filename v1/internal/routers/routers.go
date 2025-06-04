package routers

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/maneulf/guarapo_lab_test/internal/handlers"
	apmiddlewares "github.com/maneulf/guarapo_lab_test/internal/middlewares"
	"github.com/maneulf/guarapo_lab_test/internal/services"
)

// Router ...
type Router struct {
	Eng          *gin.Engine
	H            *handlers.Handlers
	M            *apmiddlewares.AuthMiddleware
	TasksService services.TasksService
}

// InitRouters set endpoints with yours handlers
func (r *Router) InitRouters() {
	// init v1 routers
	r.initV1Routers()
}

// initV1Routers set endpoints with yours handlers
func (r *Router) initV1Routers() {

	r.Eng.POST("/api/v1/login", r.H.Login)
	apiv1 := r.Eng.Group("/api/v1")

	// middleware for authorization
	apiv1.Use(r.M.Auth)
	{
		apiv1.GET("/tasks", r.H.GetTasks)
		apiv1.GET("/tasks/:id", r.H.GetTask)
		apiv1.POST("/tasks", r.H.SaveTask)
		apiv1.PUT("/tasks/:id", r.H.UpdateTask)
		apiv1.DELETE("/tasks/:id", r.H.DeleteTask)
	}
}
