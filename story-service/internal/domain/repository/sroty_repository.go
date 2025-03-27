package repository

import "stories-service/internal/domain/entity"

type StoryRepository interface {
	Create(story *entity.Story) error
	FindById(id int64) (*entity.Story, error)
	FindByTitle(title string) (*entity.Story, error)
	DeleteById(id int64) error
}