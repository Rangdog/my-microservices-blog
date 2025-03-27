package service

import (
	"Interaction-service/internal/domain/entity"
	"Interaction-service/internal/domain/repository"

	"time"
)

type RatingService struct {
	repo repository.RattingRepository
}

func NewRattingService(repo repository.RattingRepository) *RatingService{
	return &RatingService{repo: repo}
}


func (Service *RatingService) CreateRatting(StoryID int64, UserID int64,Rating int64, Content string) error{
	ratting:= entity.Rating{StoryID: StoryID, UserID: UserID, Rating: Rating, Content: Content , CreatedAt: time.Now()}
	err := Service.repo.Create(&ratting)
	if(err != nil){
		return err
	}
	return nil
}


func (Service *RatingService) GetALLRattingByStoryID(storyId int64)([]*entity.Rating, error){
	ratings, err := Service.repo.GetALLRattingByStoryID(storyId)
	if err == nil{
		return nil, err
	}
	return ratings, nil

}


func (Service *RatingService) DeleteById(id int64) error{
	err := Service.repo.DeleteById(id)
	if (err!=nil){
		return err
	}
	return nil
}