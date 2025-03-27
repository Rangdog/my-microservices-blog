package service

import (
	"Interaction-service/internal/domain/entity"
	"Interaction-service/internal/domain/repository"

	"time"
)

type FolowService struct {
	repo repository.FollowRepository
}

func NewFolowService(repo repository.FollowRepository) *FolowService{
	return &FolowService{repo: repo}
}


func (Service *FolowService) CreateFolow(StoryID int64, UserID int64) error{
	follow:= entity.Follow{StoryID: StoryID, UserID: UserID, CreatedAt: time.Now()}
	err := Service.repo.Create(&follow)
	if(err != nil){
		return err
	}
	return nil
}


func (Service *FolowService) GetALLFollowByStoryID(storyId int64)([]*entity.Follow, error){
	follow, err := Service.repo.GetALLFolowByStoryID(storyId)
	if err == nil{
		return nil, err
	}
	return follow, nil

}


func (Service *FolowService) GetALLFollowByUserID(userId int64)([]*entity.Follow, error){
	follow, err := Service.repo.GetALLFolowByUserID(userId)
	if err == nil{
		return nil, err
	}
	return follow, nil

}

func (Service *FolowService) DeleteById(id int64) error{
	err := Service.repo.DeleteById(id)
	if (err!=nil){
		return err
	}
	return nil
}