package repository

import (
	"Interaction-service/internal/domain/entity"
)

type FavoriteRepository interface {
	Create(story *entity.Favorite) error
	GetALLFavoriteByStoryID(storyId int64) ([]*entity.Favorite, error)
	GetALLFavoriteByUserID(userId int64) ([]*entity.Favorite, error)
	DeleteById(id int64) error
}