package database

import (
	"Interaction-service/internal/domain/entity"
	"Interaction-service/internal/domain/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type MySQLFavoriteRepository struct {
	db *gorm.DB
}

func NewMySQLFavoriteRepository(db *gorm.DB) repository.FavoriteRepository{
	return &MySQLFavoriteRepository{db: db}
}

func(r *MySQLFavoriteRepository) Create(favorite *entity.Favorite) error{
	if err:=r.db.Create(favorite).Error; err != nil{
		return err
	}
	return nil
}

func(r *MySQLFavoriteRepository) GetALLFavoriteByStoryID(storyID int64) ([]*entity.Favorite, error){
	var favorites []*entity.Favorite
	result:=r.db.Where("story_id = ?", storyID).Find(&favorites)
	if result.Error != nil{
		return nil, result.Error
	}
	return favorites, nil
}

func(r *MySQLFavoriteRepository) GetALLFavoriteByUserID(userId int64) ([]*entity.Favorite, error){
	var favorites []*entity.Favorite
	result:=r.db.Where("story_id = ?", userId).Find(&favorites)
	if result.Error != nil{
		return nil, result.Error
	}
	return favorites, nil
}


func (r *MySQLFavoriteRepository) DeleteById(favoriteID int64) error{
	result := r.db.Model(&entity.Favorite{}).Where("id = ? AND delete_at IS NULL", favoriteID).Update("delete_at" ,time.Now())
	if result.Error != nil{
		return result.Error
	}
	if result.RowsAffected == 0{
		return errors.New("comment not found or already deleted")
	}
	return nil
}