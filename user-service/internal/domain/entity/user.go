package entity
import "time"

type User struct{
	ID int64 `json:"id"`
	Email string `json:"email"`
	Password string	`json:"-"` // Không trả về password khi response
	CreatedAt time.Time `json:"created_at"`
}