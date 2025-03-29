package entity

import (
	"time"

	"gorm.io/gorm"
)

type Story struct {
	ID          int64  `gorm:"primaryKey; autoIncrement" json:"id"`
	Title       string `gorm:"type:varchar(255); not null" json:"title"`
	ImageStory  []byte	`gorm:"type:mediumblob" json:"image_story"`
	Description *string `gorm:"type:text" json:"description"`
	Author_id   int64  `gorm:"not nul" json:"author_id"`
	Status      string `gorm:"type:enum('finish','continue','pause');default:'finish'" json:"status"`
	Created_at  time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at  time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP; autoUpdateTime" json:"updated_at"`
	Deleted_at  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}