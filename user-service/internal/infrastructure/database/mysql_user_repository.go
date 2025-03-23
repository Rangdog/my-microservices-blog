package database

import (
	"database/sql"
	"errors"
	"user-service/internal/domain/entity"
	"user-service/internal/domain/repository"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) repository.UserRepository {
	return &MySQLUserRepository{db: db}
}

func(r *MySQLUserRepository) Create(user *entity.User) error{
	query := "INSERT INTO users(email, password, created_at) VALUES(?,?,?)"
	result, err := r.db.Exec(query,user.Email, user.Password, user.CreatedAt)
	if(err != nil){
		return err
	}

	id, err:= result.LastInsertId()
	if err != nil{
		return err
	}
	user.ID = int64(id)
	return nil
}

func(r *MySQLUserRepository) FindByEmail(email string) (*entity.User, error){
	user := &entity.User{}
	query := "SELECT id, email, password, created_at FROM users WHERE email = ?"
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err == sql.ErrNoRows{
		return nil, errors.New("User not found")
	} 
	if err != nil{
		return nil, err
	}
	return user, nil
}