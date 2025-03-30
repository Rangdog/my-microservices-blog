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

func (h *GatewayHandler) ProxyToService(serviceName string) gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Printf("[DEBUG] Gateway received - Method: %s, Path: %s, Headers: %v\n", c.Request.Method, c.Request.URL.Path, c.Request.Header)
        services, err := h.consulClient.DiscoverService(serviceName)
        if err != nil {
            c.JSON(http.StatusServiceUnavailable, gin.H{"error": fmt.Sprintf("Service %s not available", serviceName)})
            return
        }
        service := services[0]
        ip := service.Service.Address
        port := 80
        targetURL := fmt.Sprintf("http://%s.default.svc.cluster.local:%d%s", ip, port, c.Request.URL.Path)
        fmt.Printf("[DEBUG] Proxying to: %s\n", targetURL)
        url, err := url.Parse(targetURL)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse target URL"})
            return
        }
        proxy := httputil.NewSingleHostReverseProxy(url)
        // Ghi lại lỗi hoặc phản hồi từ proxy
        proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
            fmt.Printf("[ERROR] Proxy failed: %v\n", err)
            c.JSON(http.StatusBadGateway, gin.H{"error": "Bad Gateway", "details": err.Error()})
        }
        // Ghi lại phản hồi từ user-service
        originalWriter := c.Writer
        c.Writer = &responseLogger{ResponseWriter: originalWriter}
        proxy.ServeHTTP(c.Writer, c.Request)
    }
}

// Struct để ghi log phản hồi
type responseLogger struct {
    gin.ResponseWriter
}

func (w *responseLogger) Write(b []byte) (int, error) {
    fmt.Printf("[DEBUG] Response from user-service: %s\n", string(b))
    return w.ResponseWriter.Write(b)
}

func (h *GatewayHandler) HealthCheck(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}