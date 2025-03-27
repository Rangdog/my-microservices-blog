package repository

import "stories-service/internal/domain/entity"

type ChapterRepository interface {
	Create(chapter *entity.Chapter) error
	FindByStoryId(storyId int64) (*entity.Chapter, error)
}