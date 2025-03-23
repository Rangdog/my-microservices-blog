package service

import (
	"errors"
	"time"
	"user-service/internal/domain/entity"
	"user-service/internal/domain/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
	jwtSecret string
}

func NewUserService(repo repository.UserRepository, jwtSecret string) *UserService{
	return &UserService{repo: repo, jwtSecret: jwtSecret}
}

func (Service *UserService) Register(email, password string)(*entity.User, error){
	if _, err := Service.repo.FindByEmail(email); err == nil{
		return nil, errors.New(("Email already exists"))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return nil,err
	}

	user := &entity.User{
		Email: email,
		Password: string(hashedPassword),
		CreatedAt: time.Now(),
	}

	if err := Service.repo.Create(user); err != nil{
		return nil, err
	}
	return user, nil

}

func (Service *UserService) Login(email, password string)(*entity.User,string,error){
	user,err:=Service.repo.FindByEmail(email)
	if err != nil{
		return nil, "", errors.New("Invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil{
		return nil,"", errors.New(("Invalid email or password"))
	}

	claims:= jwt.MapClaims{
		"user_id": user.ID,
		"email": user.Email,
		"exp": time.Now().Add(time.Hour*24).Unix(),
	}

	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(Service.jwtSecret))
	
	if err != nil{
		return nil, "", err
	}


	return user, tokenString ,nil
}