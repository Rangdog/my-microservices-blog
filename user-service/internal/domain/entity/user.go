package entity

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"` 
	CreatedAt time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_at"`
	Role      string         `gorm:"type:enum('user', 'admin');default:'user'" json:"role"`
}