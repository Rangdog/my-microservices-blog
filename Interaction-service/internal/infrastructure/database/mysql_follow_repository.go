package database

import (
	"Interaction-service/internal/domain/entity"
	"Interaction-service/internal/domain/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type MySQLFolowRepository struct {
	db *gorm.DB
}

func NewMySQLFolowRepository(db *gorm.DB) repository.FollowRepository{
	return &MySQLFolowRepository{db: db}
}

func(r *MySQLFolowRepository) Create(follow *entity.Follow) error{
	if err:=r.db.Create(follow).Error; err != nil{
		return err
	}
	return nil
}

func(r *MySQLFolowRepository) GetALLFolowByStoryID(storyID int64) ([]*entity.Follow, error){
	var follows []*entity.Follow
	result:=r.db.Where("story_id = ?", storyID).Find(&follows)
	if result.Error != nil{
		return nil, result.Error
	}
	return follows, nil
}

func(r *MySQLFolowRepository) GetALLFolowByUserID(userId int64) ([]*entity.Follow, error){
	var follows []*entity.Follow
	result:=r.db.Where("story_id = ?", userId).Find(&follows)
	if result.Error != nil{
		return nil, result.Error
	}
	return follows, nil
}


func (r *MySQLFolowRepository) DeleteById(followID int64) error{
	result := r.db.Model(&entity.Favorite{}).Where("id = ? AND delete_at IS NULL", followID).Update("delete_at" ,time.Now())
	if result.Error != nil{
		return result.Error
	}
	if result.RowsAffected == 0{
		return errors.New("comment not found or already deleted")
	}
	return nil
}