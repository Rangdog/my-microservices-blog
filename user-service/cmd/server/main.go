package main

import (
	"database/sql"
	"net/http"
	"user-service/api"
	"user-service/api/handlers"
	"user-service/config"
	"user-service/internal/domain/service"
	"user-service/internal/infrastructure/database"
	logger "user-service/internal/pkg/Logger"

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

	userRepo := database.NewMySQLUserRepository(db)
	UserService := service.NewUserService(userRepo, cfg.JWTSecret)
	UserHandler := handlers.NewUserHandler(UserService)

	mux := api.SetupRoutes(UserHandler, cfg.JWTSecret)

	logger.Info("Server running on " + cfg.Port)
	if err := http.ListenAndServe(cfg.Port, mux); err != nil{
		logger.Error("Server failed", err)
	}
}