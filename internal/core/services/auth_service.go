package services

import (
	"fmt"
	"time"

	"github.com/bergsantana/edu-magic-api/internal/adapters/repositories"
	"github.com/bergsantana/edu-magic-api/internal/core/domain"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo  repositories.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo repositories.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Signup(user *domain.User) error {
	existingUser, _ := s.userRepo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}
	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = string(hashedPassword)

	return s.userRepo.CreateUser(user)
}

func (s *AuthService) Login(email, password string) (*domain.User, string, error) {
	// Check if user exists
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return nil, "", fmt.Errorf("invalid email or password")
	}

	// Validate password using bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", fmt.Errorf("invalid email or password")
	}

	token, err := s.generateJWT(user.ID)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %v", err)
	}
	return user, token, nil
}

func (s *AuthService) generateJWT(userID int64) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	})
	return token.SignedString([]byte(s.jwtSecret))
}
