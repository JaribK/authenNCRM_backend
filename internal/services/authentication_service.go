package services

import (
	"authenncrm/config"
	"authenncrm/internal/entities"
	"authenncrm/internal/repositories"
	"authenncrm/internal/utils"
	"authenncrm/pkg/authentication"
	"errors"
)

type (
	AuthenticationService interface {
		Login(request entities.LoginRequest) (entities.TokenResponse, error)
	}

	authenticationService struct {
		userRepo repositories.UserRepository
		config   *config.Config
	}
)

func NewAuthenticationService(userRepo repositories.UserRepository, config *config.Config) AuthenticationService {
	return &authenticationService{
		userRepo: userRepo,
		config:   config,
	}
}

func (s *authenticationService) Login(request entities.LoginRequest) (entities.TokenResponse, error) {
	user, err := s.userRepo.CheckUserIdentifierExists(request.Identifier)
	if err != nil {
		return entities.TokenResponse{}, err
	}

	if !utils.CheckPasswordEncryption(request.Password, user.Password) {
		return entities.TokenResponse{}, errors.New("invalid password")
	}

	accessToken, err := authentication.GenerateToken(user, s.config)
	if err != nil {
		return entities.TokenResponse{}, err
	}

	refreshToken, err := authentication.GenerateRefreshToken(user, s.config)
	if err != nil {
		return entities.TokenResponse{}, err
	}

	return entities.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
