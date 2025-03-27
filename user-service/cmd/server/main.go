package main

import (
	"context"
	"log"
	"microservices/pkg/discovery"
	"net/http"
	"user-service/api"
	"user-service/api/handlers"
	"user-service/config"
	"user-service/internal/domain/entity"
	"user-service/internal/domain/service"
	"user-service/internal/infrastructure/database"
	logger "user-service/internal/pkg/Logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	dsn:="thanh:123@tcp(localhost:3306)/user_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		logger.Error("Failed to connect to Mysql", err)
		return
	}

	if err := db.AutoMigrate(&entity.User{}); err!=nil{
		log.Fatal("Failed Migration")
	}
	logger.Info("Connected to MySQL")


	consulClient, err := discovery.NewConsulClient("consul:8500", "user-service", "user-service", 8080)
	if err != nil{
		log.Fatal("Failed to initalize Consul client: ", err)
	}

	ctx, _ := context.WithCancel(context.Background())
	go consulClient.StartHeartbeat(ctx)

	userRepo := database.NewMySQLUserRepository(db)
	UserService := service.NewUserService(userRepo, cfg.JWTSecret)
	UserHandler := handlers.NewUserHandler(UserService)

	mux := api.SetupRoutes(UserHandler, cfg.JWTSecret)

	logger.Info("Server running on " + cfg.Port)
	if err := http.ListenAndServe(cfg.Port, mux); err != nil{
		logger.Error("Server failed", err)
	}
}