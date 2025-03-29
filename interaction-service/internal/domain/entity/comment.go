package entity

import (
	"time"

	"gorm.io/gorm"
)

// Comment đại diện cho một bình luận trên story
type Comment struct {
    ID        int64            `gorm:"primaryKey;autoIncrement" json:"id"`
    StoryID   int64            `gorm:"not null" json:"story_id"` // Không phải foreign key
    UserID    int64            `gorm:"not null" json:"user_id"`  // Không phải foreign key
    Content   string         `gorm:"type:text;not null" json:"content"`
    CreatedAt time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}