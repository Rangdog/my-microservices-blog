package entity

import (
	"time"

	"gorm.io/gorm"
)

// Favorite đại diện cho việc một user thêm story vào danh sách yêu thích
type Favorite struct {
    ID        int64         `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID    int64            `gorm:"not null;uniqueIndex:idx_favorite" json:"user_id"`  // Không phải foreign key
    StoryID   int64             `gorm:"not null;uniqueIndex:idx_favorite" json:"story_id"` // Không phải foreign key
    CreatedAt time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}