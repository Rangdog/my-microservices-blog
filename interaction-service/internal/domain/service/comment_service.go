package service

import (
	"Interaction-service/internal/domain/entity"
	"Interaction-service/internal/domain/repository"

	"time"
)

type CommentService struct {
	repo repository.CommentRepository
}


func NewCommentService(repo repository.CommentRepository) *CommentService{
	return &CommentService{repo: repo}
}




func (Service *CommentService) CreateComment(StoryID int64, UserID int64, Content string) error{
	comment:= entity.Comment{StoryID: StoryID, UserID: UserID, Content: Content, CreatedAt: time.Now()}
	err := Service.repo.Create(&comment)
	if(err != nil){
		return err
	}
	return nil
}
func (Service *CommentService) GetALLCommentByStoryID(storyId int64)([]*entity.Comment, error){
	comment, err := Service.repo.GetALLCommentByStoryID(storyId)
	if err == nil{
		return nil, err
	}
	return comment, nil

}

func (Service *CommentService) DeleteById(id int64) error{
	err := Service.repo.DeleteById(id)
	if (err!=nil){
		return err
	}
	return nil
}