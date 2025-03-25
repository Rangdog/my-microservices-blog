package entity

import "time"

type Story struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description *string `json:"description"`
	Author_id   int64  `json:"author_id"`
	Status      string `json:"status"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	Deleted_at  *time.Time `json:"deleted_at"`
}