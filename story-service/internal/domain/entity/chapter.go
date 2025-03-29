package entity

import (
	"time"

	"gorm.io/gorm"
)

type Chapter struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	StoryID   *int64          `gorm:"not null;index" json:"story_id"` // ✅ Đổi thành `StoryID`
	Title     string         `gorm:"type:varchar(255);not null" json:"title"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"` // ✅ Sử dụng `autoCreateTime`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Story Story `gorm:"foreignKey:StoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"story"`
}
