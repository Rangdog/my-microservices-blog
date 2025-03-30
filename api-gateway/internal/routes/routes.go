package routes

import (
	"api-gateway/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handler.GatewayHandler){
	r.GET("/health", h.HealthCheck)

	users:=r.Group("/api/user-service")
	users.Any("/*any", h.ProxyToService("user-service"))

	stories := r.Group("/api/story-service")
	stories.Any("/*any", h.ProxyToService("story-service"))

	interaction := r.Group("/api/interaction-service")
	interaction.Any("/*any", h.ProxyToService("interaction-service"))
}