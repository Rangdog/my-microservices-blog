package main

import (
	"context"
	"log"
	"microservices/pkg/discovery"

	"api-gateway/config"
	"api-gateway/internal/handler"
	"api-gateway/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	consulClient, err := discovery.NewConsulClient("consul:8500", "api-gateway", "api-gateway", 8080)
	if err != nil{
		log.Fatal("Failed to initialize Consul client", err)
	}

	ctx, _ :=  context.WithCancel(context.Background())
	go consulClient.StartHeartbeat(ctx)

	gatewayHandler:= handler.NewGatewayHandler(consulClient)
	r:=gin.Default()
	routes.SetupRoutes(r, gatewayHandler)

	if err := r.Run(":8080"); err != nil{
		log.Fatal("failed to run server:", err)
	}
	
}