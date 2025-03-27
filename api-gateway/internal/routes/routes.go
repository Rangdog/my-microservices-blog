package routes

import (
	"api-gateway/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handler.GatewayHandler){
	r.GET("/health", h.HealthCheck)

	users:=r.Group("/api/users")
	users.Any("/*any", h.ProxyToService("user-service"))

	stories := r.Group("/api/stories")
	stories.Any("/*any", h.ProxyToService("story-service"))
}