package services

import (
	"authenncrm/internal/constants"
	"authenncrm/internal/entities"
	"authenncrm/internal/repositories"
	"authenncrm/internal/utils"
	"errors"
)

type (
	UserService interface {
		CreateUser(user *entities.UserRequest) (*entities.UserResponse, error)
		GetUserById(id string) (*entities.UserResponse, error)
		GetAllUsers() ([]entities.UserResponse, error)
		UpdateUser(id string, user *entities.UserRequest) error
		DeleteUser(id string) error
	}

	userService struct {
		userRepo repositories.UserRepository
	}
)

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(request *entities.UserRequest) (*entities.UserResponse, error) {

	if err := s.userRepo.CheckUserNameOrEmailExists(request); err != nil {
		return nil, err
	}

	if response := utils.ValidatePassword(request); response != constants.PassedValidation {
		return nil, errors.New(response)
	}

	user := &entities.User{
		Username:  request.Username,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		CreatedBy: utils.StringToUUID(request.CreatedBy),
		UpdatedBy: utils.StringPtrToUUIDPtr(request.UpdatedBy),
	}

	if request.RoleId != nil {
		roleId := utils.StringPtrToUUIDPtr(request.RoleId)
		user.RoleId = roleId
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, err
	}

	response := &entities.UserResponse{
		Id:        user.Id,
		Username:  request.Username,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Points:    0,
	}

	return response, nil
}

func (s *userService) GetUserById(id string) (*entities.UserResponse, error) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	response := &entities.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role: &entities.RoleResponse{
			Id:    user.Role.Id,
			Name:  user.Role.Name,
			Level: user.Role.Level,
		},
		Points: 0,
	}

	for _, log := range user.PointLogs {
		response.Points += log.Points
	}

	return response, nil
}

func (s *userService) GetAllUsers() ([]entities.UserResponse, error) {
	users, err := s.userRepo.GetUsers()
	if err != nil {
		return nil, err
	}

	var responses []entities.UserResponse
	for _, user := range users {
		response := entities.UserResponse{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
			Role: &entities.RoleResponse{
				Id:    user.Role.Id,
				Name:  user.Role.Name,
				Level: user.Role.Level,
			},
			Points: 0,
		}

		for _, log := range user.PointLogs {
			response.Points += log.Points
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (s *userService) UpdateUser(id string, user *entities.UserRequest) error {
	existingUser, err := s.userRepo.GetUserById(id)
	if err != nil {
		return err
	}

	existingUser.Username = user.Username
	existingUser.Password = user.Password
	existingUser.Email = user.Email

	if err := s.userRepo.UpdateUser(existingUser); err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(id string) error {
	if err := s.userRepo.DeleteUser(id); err != nil {
		return err
	}
	return nil
}
