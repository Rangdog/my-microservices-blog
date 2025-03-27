package entity

import (
	"time"

	"gorm.io/gorm"
)

// Rating đại diện cho một đánh giá của user đối với story
type Rating struct {
    ID        int64            `gorm:"primaryKey;autoIncrement" json:"id"`
    StoryID   int64            `gorm:"not null;uniqueIndex:idx_rating" json:"story_id"`
    UserID    int64            `gorm:"not null;uniqueIndex:idx_rating" json:"user_id"` // Không phải foreign key
    Rating    int64            `gorm:"check:rating >= 1 AND rating <= 5" json:"rating"`
    Content   string         `gorm:"type:text" json:"content"`
    CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}