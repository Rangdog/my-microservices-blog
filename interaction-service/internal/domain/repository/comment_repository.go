package repository

import (
	"Interaction-service/internal/domain/entity"
)

type CommentRepository interface {
	Create(comment *entity.Comment) error
	GetALLCommentByStoryID(storyId int64) ([]*entity.Comment, error)
	DeleteById(id int64) error
}