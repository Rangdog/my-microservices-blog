package repository

import (
	"Interaction-service/internal/domain/entity"
)

type FollowRepository interface {
	Create(story *entity.Follow) error
	GetALLFolowByStoryID(storyId int64) ([]*entity.Follow, error)
	GetALLFolowByUserID(userId int64) ([]*entity.Follow, error)
	DeleteById(id int64) error
}