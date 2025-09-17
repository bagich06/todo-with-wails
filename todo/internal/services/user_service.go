package services

import (
	"errors"
	"todo/internal/models"
	"todo/internal/repository"
)

type UserService struct {
	repo       repository.UserRepositoryInterface
	jwtService *JWTService
}

func NewUserService(repo repository.UserRepositoryInterface, jwtService *JWTService) *UserService {
	return &UserService{repo: repo, jwtService: jwtService}
}

func (s *UserService) Login(email, password string) (*models.AuthResponse, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("wrong password")
	}

	token, err := s.jwtService.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

func (s *UserService) Register(username, email, password string) (*models.AuthResponse, error) {
	exists, err := s.repo.UserExists(email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("user with this email already exists")
	}

	var user models.User
	user = models.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	userID, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, errors.New("error while creating user")
	}

	user.ID = userID
	user.Password = ""

	token, err := s.jwtService.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	authResponse := &models.AuthResponse{
		Token: token,
		User:  user,
	}

	return authResponse, nil
}
