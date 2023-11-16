package repository

import (

	"architecture.com/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(user *domain.User) error {
	err := ur.db.Create(user).Error
	return err
}

func (ur *userRepository) Fetch() ([]domain.User, error) {
	var users []domain.User
	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (ur *userRepository) GetByID(id string) (domain.User, error) {
	var user domain.User
	err := ur.db.First(&user, "id = ?", id).Error
	return user, err
}