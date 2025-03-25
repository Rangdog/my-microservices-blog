package main

import (
	"context"
	"database/sql"
	"log"
	"microservices/pkg/discovery"
	"net/http"
	"stories-service/api"
	"stories-service/api/handlers"
	"stories-service/config"
	"stories-service/internal/domain/service"
	"stories-service/internal/infrastructure/database"
	logger "stories-service/internal/pkg/Logger"

	_ "github.com/go-sql-driver/mysql"
)

// dấu _ trong imort là giúp go chỉ tải driver mà không dùng trực tiếp
func main() {
	cfg := config.LoadConfig()
	if cfg.Port == ""{
		cfg.Port = ":8080"
	}

	if  cfg.JWTSecret == ""{
		logger.Error("JWT_SECRET is required", nil)
		return
	}

	db, err := sql.Open("mysql", cfg.MySQLDSN)
	if err != nil{
		logger.Error("Failed to connect to Mysql", err)
		return
	}

	defer db.Close()

	if err := db.Ping(); err != nil{
		logger.Error("MySQL ping failed", err)
		return
	}
	logger.Info("Connected to MySQL")

	consulClient, err := discovery.NewConsulClient("consul:8500", "story-service", "story-service", 8100)
	if err != nil{
		log.Fatal("Failed to initialize Consul client", err)
	}
	ctx,_:=context.WithCancel(context.Background())
	go consulClient.StartHeartbeat(ctx)

	userRepo := database.NewMySQLStoryRepository(db)
	UserService := service.NewStoryService(userRepo)
	UserHandler := handlers.NewStoryHandler(UserService)

	mux := api.SetupRoutes(UserHandler, cfg.JWTSecret)

	logger.Info("Server running on " + cfg.Port)
	if err := http.ListenAndServe(cfg.Port, mux); err != nil{
		logger.Error("Server failed", err)
	}
}