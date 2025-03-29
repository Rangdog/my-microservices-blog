package main

import (
	"Interaction-service/api"
	"Interaction-service/api/handlers"
	"Interaction-service/config"
	"Interaction-service/internal/domain/entity"
	"Interaction-service/internal/domain/service"
	"Interaction-service/internal/infrastructure/database"
	logger "Interaction-service/internal/pkg/Logger"
	"context"
	"log"
	"microservices/pkg/discovery"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// dấu _ trong imort là giúp go chỉ tải driver mà không dùng trực tiếp
func main() {
	cfg := config.LoadConfig()
	if cfg.Port == ""{
		cfg.Port = ":8200"
	}

	dsn := "thanh:123@tcp(mysql:3306)/interaction_db?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		logger.Error("Failed to connect to Mysql", err)
		return
	}
	// Tự động migrate các bảng
    if err := db.AutoMigrate(
        &entity.Comment{},
        &entity.Rating{},
        &entity.Favorite{},
        &entity.Follow{},
    ); err != nil {
        log.Fatal("failed Migration:", err)
    }
	logger.Info("Connected to MySQL")

	consulClient, err := discovery.NewConsulClient("consul:8500", "story-service", "story-service", 8100)
	if err != nil{
		log.Fatal("Failed to initialize Consul client", err)
	}
	ctx,_:=context.WithCancel(context.Background())
	go consulClient.StartHeartbeat(ctx)

	//comment
	CommentRepo := database.NewMySQLCommentRepository(db)
	CommentService:= service.NewCommentService(CommentRepo)
	CommentHandler:= handlers.NewCommentHandler(CommentService) 

	//favotite

	FavoriteRepo := database.NewMySQLFavoriteRepository(db)
	FavoriteService:= service.NewFavoriteService(FavoriteRepo)
	FavoriteHandler:= handlers.NewFavoriteHandler(FavoriteService) 

	//follow

	FollowRepo := database.NewMySQLFolowRepository(db)
	FollowService:= service.NewFolowService(FollowRepo)
	FollowHandler:= handlers.NewFollowHandler(FollowService) 

	//rating

	RatingRepo := database.NewMySQLRattingRepository(db)
	RatingService:= service.NewRattingService(RatingRepo)
	RatingHandler:= handlers.NewRatingHandler(RatingService) 

	mux := api.SetupRoutes(CommentHandler, FavoriteHandler, FollowHandler, RatingHandler)

	logger.Info("Server running on " + cfg.Port)
	if err := http.ListenAndServe(cfg.Port, mux); err != nil{
		logger.Error("Server failed", err)
	}
}