package services

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"bookstore/backend/internal/models"
	"bookstore/backend/internal/repository"
)

type AuthService struct {
	repo      *repository.UserRepository
	jwtSecret string
}

func NewAuthService(repo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Register(ctx context.Context, req models.RegisterRequest) (*models.AuthResponse, error) {
	req.FullName = strings.TrimSpace(req.FullName)
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Password = strings.TrimSpace(req.Password)

	if req.FullName == "" || req.Email == "" || req.Password == "" {
		return nil, errors.New("all fields are required")
	}

	existing, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("user already exists")
	}

	user, err := s.repo.Create(ctx, req.FullName, req.Email, req.Password, "user")
	if err != nil {
		return nil, err
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = ""
	return &models.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error) {
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Password = strings.TrimSpace(req.Password)

	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	if user.PasswordHash != req.Password {
		return nil, errors.New("invalid credentials")
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = ""
	return &models.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *AuthService) GetByID(ctx context.Context, id int64) (*models.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user != nil {
		user.PasswordHash = ""
	}
	return user, nil
}

func (s *AuthService) ParseToken(tokenString string) (int64, string, error) {
	decoded, err := base64.StdEncoding.DecodeString(tokenString)
	if err != nil {
		return 0, "", errors.New("invalid token")
	}

	parts := strings.Split(string(decoded), ":")
	if len(parts) != 3 {
		return 0, "", errors.New("invalid token format")
	}

	userID, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, "", errors.New("invalid user id")
	}

	role := parts[1]
	secret := parts[2]

	if secret != s.jwtSecret {
		return 0, "", errors.New("invalid token secret")
	}

	return userID, role, nil
}

func (s *AuthService) generateToken(user *models.User) (string, error) {
	raw := fmt.Sprintf("%d:%s:%s", user.ID, user.Role, s.jwtSecret)
	return base64.StdEncoding.EncodeToString([]byte(raw)), nil
}
