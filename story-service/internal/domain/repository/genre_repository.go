package repository

import "stories-service/internal/domain/entity"

type GenreRepository interface { 
	Create(genre *entity.Genre) error
	FindByName(name string) (*entity.Genre, error)
}