package repository

import "stories-service/internal/domain/entity"

type StoryGenreRepository interface {
	Create(StoryGenre *entity.StoryGenre) error
	DeleteById(storyId int64, genreId int64) error
}