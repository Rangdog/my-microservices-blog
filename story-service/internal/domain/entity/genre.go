package entity

type Genre struct {
	ID   int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(50);unique;not null" json:"name"`
}