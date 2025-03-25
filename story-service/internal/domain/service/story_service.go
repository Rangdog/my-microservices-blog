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
	if err == nil{
		return nil, err
	}
	return story, nil

}

func (Service *StoryService) CreateStory(title string, description string, author_id int64) error{
	story:=entity.Story{Title:title,Description:&description,Author_id:author_id,Status:"continue",Created_at:time.Now(),Updated_at:time.Now()}
	err := Service.repo.Create(&story)
	if(err != nil){
		return err
	}
	return nil
}

func (Service *StoryService) DeleteById(id int64) error{
	err := Service.repo.DeleteById(id)
	if (err!=nil){
		return err
	}
	return nil
}