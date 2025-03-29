package main

import (
	"context"
	"log"
	"microservices/pkg/discovery"
	"net/http"
	"stories-service/api"
	"stories-service/api/handlers"
	"stories-service/config"
	"stories-service/internal/domain/entity"
	"stories-service/internal/domain/service"
	"stories-service/internal/infrastructure/database"
	logger "stories-service/internal/pkg/Logger"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// dấu _ trong imort là giúp go chỉ tải driver mà không dùng trực tiếp
func main() {
	cfg := config.LoadConfig()
	if cfg.Port == ""{
		cfg.Port = ":8100"
	}

	//Tạm thời chưa dùng đến, có thể sẽ chỉ check jwt bên API gateway
	// if  cfg.JWTSecret == ""{
	// 	logger.Error("JWT_SECRET is required", nil)
	// 	return
	// }

	dsn := "thanh:123@tcp(mysql:3306)/story_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		logger.Error("Failed to connect to Mysql", err)
		return
	}

	// Tự động migrate các bảng
    if err := db.AutoMigrate(
        &entity.Story{},
        &entity.Chapter{},
        &entity.Genre{},
        &entity.StoryGenre{},
    ); err != nil {
        log.Fatal("Migration thất bại:", err)
    }
	logger.Info("Connected to MySQL")

	consulClient, err := discovery.NewConsulClient("consul:8500", "story-service", "story-service", 8100)
	if err != nil{
		log.Fatal("Failed to initialize Consul client", err)
	}
	ctx,_:=context.WithCancel(context.Background())
	go consulClient.StartHeartbeat(ctx)


	//story
	StoryRepo := database.NewMySQLStoryRepository(db)
	StoryService := service.NewStoryService(StoryRepo)
	storyHandler := handlers.NewStoryHandler(StoryService)

	//genre
	GenreRepo := database.NewMySQLGenreRepository(db)
	GenreService := service.NewGenreService(GenreRepo)
	genreHandler:= handlers.NewGenreHandler(GenreService)

	//chapter
	ChapterRepo := database.NewMySQLChapterRepository(db)
	ChapterService:= service.NewChapterService(ChapterRepo)
	chapterHandler:=handlers.NewChapterHandler(ChapterService)

	//story_genre
	StoryGenreRepo := database.NewMySQLStoryGenreRepository(db)
	StoryGenreService := service.NewStoryGenreService(StoryGenreRepo)
	StoryGenreHandler := handlers.NewStoryGenreHandler(StoryGenreService)

	router := api.SetupRoutes(storyHandler, genreHandler, chapterHandler, StoryGenreHandler)

	logger.Info("Server running on " + cfg.Port)
	if err := http.ListenAndServe(cfg.Port, router ); err != nil{
		logger.Error("Server failed", err)
	}
}