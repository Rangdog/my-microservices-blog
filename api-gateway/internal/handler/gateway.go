package handler

import (
	"fmt"
	"microservices/pkg/discovery"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type GatewayHandler struct {
	consulClient *discovery.ConsulClient
}

func NewGatewayHandler(consulClient *discovery.ConsulClient) *GatewayHandler{
	return &GatewayHandler{consulClient: consulClient}
}

func (h *GatewayHandler) ProxyToService(serviceName string) gin.HandlerFunc{
	return func(c *gin.Context){
		services, err := h.consulClient.DiscoverService(serviceName)
		if err != nil{
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": fmt.Sprintf("Service %s not available", serviceName)})
			return
		}

		service := services[0]
		targetURL := fmt.Sprintf("http://%s:%d", service.Service.Address, service.Service.Port)
		
		url,err := url.Parse(targetURL)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse target URL"})
			return
		}

		proxy:=httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(c.Writer, c.Request)

	}
}

func (h *GatewayHandler) HealthCheck(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}