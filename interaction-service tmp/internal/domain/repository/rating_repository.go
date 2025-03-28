package repository

import (
	"Interaction-service/internal/domain/entity"
)

type RattingRepository interface {
	Create(rating *entity.Rating) error
	GetALLRattingByStoryID(storyID int64) ([]*entity.Rating, error)
	DeleteById(id int64) error
}