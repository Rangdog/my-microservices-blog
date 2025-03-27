package database

import (
	"errors"
	"stories-service/internal/domain/entity"
	"stories-service/internal/domain/repository"

	"gorm.io/gorm"
)

type MySQLGenreRepository struct {
	db *gorm.DB
}

func NewMySQLGenreRepository(db *gorm.DB) repository.GenreRepository{
	return &MySQLGenreRepository{db: db}
}

func(r *MySQLGenreRepository) Create(genre *entity.Genre) error{
	if err:=r.db.Create(genre).Error; err != nil{
		return err
	}
	return nil
}

func(r *MySQLGenreRepository) FindByName(name string) (*entity.Genre,error){
	var genre entity.Genre
	result:=r.db.Where("name = ?", name).First(&genre)
	if errors.Is(result.Error, gorm.ErrRecordNotFound){
		return nil, errors.New("genre not found")
	}
	if result.Error != nil{
		return nil, result.Error
	}
	return &genre, nil
}