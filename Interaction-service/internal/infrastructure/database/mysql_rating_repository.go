package database

import (
	"Interaction-service/internal/domain/entity"
	"Interaction-service/internal/domain/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type MySQLRattingRepository struct {
	db *gorm.DB
}

func NewMySQLRattingRepository(db *gorm.DB) repository.RattingRepository{
	return &MySQLRattingRepository{db: db}
}

func(r *MySQLRattingRepository) Create(rating *entity.Rating) error{
	if err:=r.db.Create(rating).Error; err != nil{
		return err
	}
	return nil
}

func(r *MySQLRattingRepository) GetALLRattingByStoryID(storyID int64) ([]*entity.Rating, error){
	var rattings []*entity.Rating
	result:=r.db.Where("story_id = ?", storyID).Find(&rattings)
	if result.Error != nil{
		return nil, result.Error
	}
	return rattings, nil
}



func (r *MySQLRattingRepository) DeleteById(rattingID int64) error{
	result := r.db.Model(&entity.Favorite{}).Where("id = ? AND delete_at IS NULL", rattingID).Update("delete_at" ,time.Now())
	if result.Error != nil{
		return result.Error
	}
	if result.RowsAffected == 0{
		return errors.New("comment not found or already deleted")
	}
	return nil
}