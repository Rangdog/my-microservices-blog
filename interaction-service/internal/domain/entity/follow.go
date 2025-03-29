package entity

import (
    "time"
    "gorm.io/gorm"
)

// Follow đại diện cho việc một user theo dõi một story
type Follow struct {
    UserID    int64            `gorm:"primaryKey" json:"user_id"`  // Không phải foreign key
    StoryID   int64            `gorm:"primaryKey" json:"story_id"` // Không phải foreign key
    CreatedAt time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}