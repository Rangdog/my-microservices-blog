package database

import (
	"database/sql"
	"errors"
	"stories-service/internal/domain/entity"
	"stories-service/internal/domain/repository"
	"time"
)

type MySQLStoryRepository struct {
	db *sql.DB
}

func NewMySQLStoryRepository(db *sql.DB) repository.StoryRepository {
	return &MySQLStoryRepository{db: db}
}

func(r *MySQLStoryRepository) Create(story *entity.Story) error{
	query := "INSERT INTO stories(title, description, author_id, status, created_at, updated_at) VALUES(?,?,?,?,?,?)"
	result, err := r.db.Exec(query,story.Title, story.Description, story.Author_id, story.Status, story.Created_at, story.Updated_at)
	if(err != nil){
		return err
	}

	id, err:= result.LastInsertId()
	if err != nil{
		return err
	}
	story.ID = int64(id)
	return nil
}

func(r *MySQLStoryRepository) FindById(id int64) (*entity.Story, error){
	story := &entity.Story{}
	query := "SELECT id, title, description, author_id, status, created_at, updated_at, deleted_at FROM stories WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&story.ID, &story.Title, &story.Description, &story.Author_id, &story.Status, &story.Created_at, &story.Updated_at, &story.Deleted_at)
	if err == sql.ErrNoRows{
		return nil, errors.New("story not found")
	} 
	if err != nil{
		return nil, err
	}
	return story, nil
}

func (r *MySQLStoryRepository) FindByTitle(title string) (*entity.Story, error){
	return &entity.Story{},nil
}

func (r *MySQLStoryRepository) DeleteById(id int64) error{
	query := "UPDATE stories SET delete_at = ? WHERE id = ? AND delete_at IS NULL"
	now:= time.Now()
	result, err := r.db.Exec(query,now, id)
	if err != nil{
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil{
		return err
	}

	if rowAffected == 0{
		return errors.New("story not found or already deleted")
	}
	return nil
}