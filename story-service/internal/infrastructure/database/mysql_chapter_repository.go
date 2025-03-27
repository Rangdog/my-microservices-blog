package database

import (
	"errors"
	"stories-service/internal/domain/entity"
	"stories-service/internal/domain/repository"

	"gorm.io/gorm"
)

type MySQLChapterRepository struct {
	db *gorm.DB
}

func NewMySQLChapterRepository(db *gorm.DB) repository.ChapterRepository{
	return &MySQLChapterRepository{db: db}
}

func(r *MySQLChapterRepository) Create(chapter *entity.Chapter) error{
	if err:=r.db.Create(chapter).Error; err != nil{
		return err
	}
	return nil
}

func(r *MySQLChapterRepository) FindByStoryId(storyID int64) (*entity.Chapter, error){
	var chapter entity.Chapter
	result:=r.db.Where("story_id = ?", storyID).First(&chapter)
	if errors.Is(result.Error, gorm.ErrRecordNotFound){
		return nil, errors.New("genre not found")
	}
	if result.Error != nil{
		return nil, result.Error
	}
	return &chapter, nil
}