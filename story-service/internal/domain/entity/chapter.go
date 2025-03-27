package entity

import (
	"time"

	"gorm.io/gorm"
)

type Chapter struct {
	ID        int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Story_id  int64 `gorm:"not null;index" json:"story_id"`
	Title     string `gorm:"type:varchar(255);not null" json:"title"`
	Created_at time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Deleted_at gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Story     Story `gorm:"foreignKey:StoryID;constraint:OnDelete:CASCADE"`
}