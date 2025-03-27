package database

import (
	"errors"
	"stories-service/internal/domain/entity"
	"stories-service/internal/domain/repository"

	"gorm.io/gorm"
)

type MySQLStoryGenreRepository struct {
	db *gorm.DB
}

func NewMySQLStoryGenreRepository(db *gorm.DB) repository.StoryGenreRepository {
	return &MySQLStoryGenreRepository{db: db}
}

func(r *MySQLStoryGenreRepository) Create(storyGenre *entity.StoryGenre) error{
	if err:=r.db.Create(storyGenre).Error; err != nil{
		return err
	}
	return nil
}


func (r *MySQLStoryGenreRepository) DeleteById(storyId int64, genreId int64) error{
	result := r.db.Where("story_id = ? AND genre_id = ?", storyId, genreId).Delete(&entity.StoryGenre{})
    
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("story-genre relation not found")
    }
    return nil
}