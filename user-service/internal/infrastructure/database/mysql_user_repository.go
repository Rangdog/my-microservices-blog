package database

import (
	"errors"
	"user-service/internal/domain/entity"
	"user-service/internal/domain/repository"

	"gorm.io/gorm"
)

type MySQLUserRepository struct {
	db *gorm.DB
}

func NewMySQLUserRepository(db *gorm.DB) repository.UserRepository {
	return &MySQLUserRepository{db: db}
}

func(r *MySQLUserRepository) Create(user *entity.User) error{
	if err:= r.db.Create(user).Error; err != nil{
		return err
	}
	return nil
}

func(r *MySQLUserRepository) FindByEmail(email string) (*entity.User, error){
	var user entity.User
	result:=r.db.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound){
		return nil, errors.New("user not found")
	}
	if result.Error != nil{
		return nil, result.Error
	}
	return &user, nil
}