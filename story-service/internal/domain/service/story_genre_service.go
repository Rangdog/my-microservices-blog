package service

import (
	"stories-service/internal/domain/entity"
	"stories-service/internal/domain/repository"
)

type StoryGenreService struct {
	repo repository.StoryGenreRepository
}

func NewStoryGenreService(repo repository.StoryGenreRepository) *StoryGenreService{
	return &StoryGenreService{repo: repo}
}
func (Service *StoryGenreService) CreateStoryGenre(storyId int64, genreId int64) error{
	storyGenre:=entity.StoryGenre{StoryID:storyId, GenreID:genreId}
	err := Service.repo.Create(&storyGenre)
	return err
}

func (Service *StoryGenreService) DeleteById(storyId int64, genreId int64) error{
	err := Service.repo.DeleteById(storyId, genreId)
	return err
}