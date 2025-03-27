package service

import (
	"stories-service/internal/domain/entity"
	"stories-service/internal/domain/repository"
)

type GenreService struct {
	repo repository.GenreRepository
}

func NewGenreService(repo repository.GenreRepository) *GenreService{
	return &GenreService{repo: repo}
}

func(Service *GenreService) CreateGenre(name string) error{
	genre := entity.Genre{Name: name}
	err:=Service.repo.Create(&genre)
	return err
}

func(Service *GenreService) FindGenreById(name string) (*entity.Genre, error){
	genre, err:=Service.repo.FindByName(name)
	return genre, err
}