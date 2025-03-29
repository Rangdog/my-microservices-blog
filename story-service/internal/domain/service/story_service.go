package service

import (
	"stories-service/internal/domain/entity"
	"stories-service/internal/domain/repository"
	"time"
)

type StoryService struct {
	repo repository.StoryRepository
}

func NewStoryService(repo repository.StoryRepository) *StoryService{
	return &StoryService{repo: repo}
}

func (Service *StoryService) GetStoryByID(id int64)(*entity.Story, error){
	story, err := Service.repo.FindById(id)
	return story,err
}

func (Service *StoryService) CreateStory(title string, description string, author_id int64) error{
	story:=entity.Story{Title:title,Description:&description,AuthorID:author_id,Status:"continue",CreatedAt:time.Now(),UpdatedAt:time.Now()}
	err := Service.repo.Create(&story)
	return err
}

func (Service *StoryService) DeleteById(id int64) error{
	err := Service.repo.DeleteById(id)
	return err
}