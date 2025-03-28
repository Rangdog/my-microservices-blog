package service

import (
	"Interaction-service/internal/domain/entity"
	"Interaction-service/internal/domain/repository"

	"time"
)

type FavoriteService struct {
	repo repository.FavoriteRepository
}

func NewFavoriteService(repo repository.FavoriteRepository) *FavoriteService{
	return &FavoriteService{repo: repo}
}


func (Service *FavoriteService) CreateFavorite(StoryID int64, UserID int64) error{
	favorite:= entity.Favorite{StoryID: StoryID, UserID: UserID, CreatedAt: time.Now()}
	err := Service.repo.Create(&favorite)
	if(err != nil){
		return err
	}
	return nil
}


func (Service *FavoriteService) GetALLFavoriteByStoryID(storyId int64)([]*entity.Favorite, error){
	favorite, err := Service.repo.GetALLFavoriteByStoryID(storyId)
	if err == nil{
		return nil, err
	}
	return favorite, nil

}


func (Service *FavoriteService) GetALLFavoriteByUserID(userId int64)([]*entity.Favorite, error){
	favorite, err := Service.repo.GetALLFavoriteByUserID(userId)
	if err == nil{
		return nil, err
	}
	return favorite, nil

}


func (Service *FavoriteService) DeleteById(id int64) error{
	err := Service.repo.DeleteById(id)
	if (err!=nil){
		return err
	}
	return nil
}