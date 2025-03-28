package database

import (
	"Interaction-service/internal/domain/entity"
	"Interaction-service/internal/domain/repository"
	"errors"
	"time"

	"gorm.io/gorm"
)

type MySQLCommentRepository struct {
	db *gorm.DB
}

func NewMySQLCommentRepository(db *gorm.DB) repository.CommentRepository{
	return &MySQLCommentRepository{db: db}
}

func(r *MySQLCommentRepository) Create(comment *entity.Comment) error{
	if err:=r.db.Create(comment).Error; err != nil{
		return err
	}
	return nil
}

func(r *MySQLCommentRepository) GetALLCommentByStoryID(storyID int64) ([]*entity.Comment, error){
	var comments []*entity.Comment
	result:=r.db.Where("story_id = ?", storyID).Find(&comments)
	if result.Error != nil{
		return nil, result.Error
	}
	return comments, nil
}


func (r *MySQLCommentRepository) DeleteById(commentId int64) error{
	result := r.db.Model(&entity.Comment{}).Where("id = ? AND delete_at IS NULL", commentId).Update("delete_at" ,time.Now())
	if result.Error != nil{
		return result.Error
	}
	if result.RowsAffected == 0{
		return errors.New("comment not found or already deleted")
	}
	return nil
}