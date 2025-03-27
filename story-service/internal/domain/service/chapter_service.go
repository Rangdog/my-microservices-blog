package service

import (
	"stories-service/internal/domain/entity"
	"stories-service/internal/domain/repository"
	"time"
)

type ChapterService struct {
	repo repository.ChapterRepository
}

func NewChapterService(repo repository.ChapterRepository) *ChapterService{
	return &ChapterService{repo:  repo}
}

func (Service *ChapterService) CreateChapter(storyId int64, title string) error{
	chapter := entity.Chapter{Story_id: storyId, Title: title, Created_at: time.Now()}
	err:= Service.repo.Create(&chapter)
	return err
}

func (Service *ChapterService) FindByStoryID(storyId int64) (*entity.Chapter, error){
	chapter,err:= Service.repo.FindByStoryId(storyId)
	return chapter,err
}