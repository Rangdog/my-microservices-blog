package repository
import "user-service/internal/domain/entity"
type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}