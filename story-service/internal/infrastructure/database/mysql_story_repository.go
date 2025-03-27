package database

import (
	"errors"
	"stories-service/internal/domain/entity"
	"stories-service/internal/domain/repository"
	"time"

	"gorm.io/gorm"
)

type MySQLStoryRepository struct {
	db *gorm.DB
}

func NewMySQLStoryRepository(db *gorm.DB) repository.StoryRepository {
	return &MySQLStoryRepository{db: db}
}

func(r *MySQLStoryRepository) Create(story *entity.Story) error{
	if err:=r.db.Create(story).Error; err != nil{
		return err
	}
	return nil
}

func(r *MySQLStoryRepository) FindById(id int64) (*entity.Story, error){
	story:=entity.Story{}
	result:=r.db.First(&story,id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound){
		return nil, errors.New("story not found")
	}
	if result.Error != nil{
		return nil, result.Error
	}
	return &story,nil
}

func (r *MySQLStoryRepository) FindByTitle(title string) (*entity.Story, error){
	return &entity.Story{},nil
}

func (r *MySQLStoryRepository) DeleteById(id int64) error{
	result := r.db.Model(&entity.Story{}).Where("id = ? AND delete_at IS NULL", id).Update("delete_at" ,time.Now())
	if result.Error != nil{
		return result.Error
	}
	if result.RowsAffected == 0{
		return errors.New("story not found or already deleted")
	}
	return nil
}