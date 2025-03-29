package entity

import (
	"time"

	"gorm.io/gorm"
)

type Story struct {
	ID          int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	ImageStory  []byte         `gorm:"type:mediumblob" json:"image_story"`
	Description *string        `gorm:"type:text" json:"description"` // ❌ Không đặt default
	AuthorID    int64          `gorm:"not null" json:"author_id"`    // ✅ Sửa lỗi `not nul`
	Status      string         `gorm:"type:enum('finish','continue','pause');default:'finish'" json:"status"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"` // ✅ Sử dụng `autoCreateTime`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"` // ✅ Sử dụng `autoUpdateTime`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Chapters []Chapter `gorm:"foreignKey:StoryID;constraint:OnDelete:CASCADE;" json:"chapters"` // ✅ Định nghĩa quan hệ ngược lại
}
