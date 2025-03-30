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
		ip := service.Service.Address  // 🛑 Lấy địa chỉ IP của user-service
		port := service.Service.Port   // 🛑 Nếu cần cổng thì lấy luôn

		// 🛑 Log địa chỉ IP để debug
		fmt.Printf("[DEBUG] Found service %s at IP: %s, Port: %d\n", serviceName, ip, port)
		// targetURL := fmt.Sprintf("http://%s:%d", service.Service.Address, service.Service.Port) bỏ port
		// targetURL := fmt.Sprintf("http://%s%s", service.Service.Address, c.Request.URL.Path) // Giữ nguyên đường dẫn gốc
		targetURL := fmt.Sprintf("http://%s.default.svc.cluster.local%s", ip, c.Request.URL.Path)
		// 🛑 Log URL để debug
		fmt.Printf("[DEBUG] Proxy request to: %s\n", targetURL)

		url,err := url.Parse(targetURL)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse target URL"})
			return
		}
		// 🛑 Log thêm thông tin về request
		fmt.Printf("[DEBUG] Request Method: %s, Path: %s, Query: %s\n", c.Request.Method, c.Request.URL.Path, c.Request.URL.RawQuery)

		proxy:=httputil.NewSingleHostReverseProxy(url)

		// 🛑 Log lỗi nếu có
		proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			fmt.Printf("[ERROR] Proxy error: %v\n", err)
			c.JSON(http.StatusBadGateway, gin.H{"error": "Bad Gateway", "details": err.Error()})
		}

		proxy.ServeHTTP(c.Writer, c.Request)

	}
}

func (h *GatewayHandler) HealthCheck(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}