package entity

type StoryGenre struct{
	StoryID int64 `gorm:"primaryKey" json:"story_id"`
	GenreID int64 `gorm:"primaryKey" json:"genre_id"`
	Story   Story `gorm:"foreignKey:StoryID;constraint:OnDelete:CASCADE"` // Foreign key trong cùng dịch vụ
    Genre   Genre `gorm:"foreignKey:GenreID;constraint:OnDelete:CASCADE"` // Foreign key trong cùng dịch vụ
}