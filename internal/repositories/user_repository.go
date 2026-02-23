package repositories

import (
	"authenncrm/internal/entities"
	"errors"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		CreateUser(user *entities.User) error
		GetUsers() ([]entities.User, error)
		GetUserById(id string) (*entities.User, error)
		GetUserByUsername(username string) (*entities.User, error)
		UpdateUser(user *entities.User) error
		DeleteUser(id string) error

		CheckUserNameOrEmailExists(request *entities.UserRequest) error
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *entities.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUsers() ([]entities.User, error) {
	var users []entities.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserById(id string) (*entities.User, error) {
	var user entities.User
	if err := r.db.
		Preload("Role").
		Preload("PointLogs").
		Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsername(username string) (*entities.User, error) {
	var user entities.User
	if err := r.db.
		Preload("Role").
		Preload("PointLogs").
		Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *entities.User) error {
	if err := r.db.Updates(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.User{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) CheckUserNameOrEmailExists(request *entities.UserRequest) error {
	type result struct {
		Username string
		Email    *string
	}

	var res result

	err := r.db.Model(&entities.User{}).
		Select("username", "email").
		Where("username = ? OR email = ?", request.Username, request.Email).
		Limit(1).
		Scan(&res).Error

	if err != nil {
		return err
	}

	if res.Username == request.Username {
		return errors.New("username already exists")
	}

	if res.Email != nil && request.Email != nil && *res.Email == *request.Email {
		return errors.New("email already exists")
	}

	return nil
}
