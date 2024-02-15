package service

import (
	"context"
	// "fmt"
	"njajal-go/config"
	errors "njajal-go/pkg"
	"njajal-go/repository"

	// entities "njajal-go/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, email string, password string) error
	Login(ctx context.Context, email string, password string) (string, error)
	ResetPassword(ctx context.Context, email string, password string) error
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (us *UserServiceImpl) Register(ctx context.Context, email string, password string) error {
	// Check if the email is already registered
	_, err := us.userRepository.FindUserByEmail(ctx, email)
	if err == nil {
		return errors.ErrUserAlreadyExists
	}
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return us.userRepository.CreateUser(ctx, email, string(hashedPassword))
}

func (us *UserServiceImpl) Login(ctx context.Context, email string, password string) (string, error) {
	// Check if the email is already registered
	user, err := us.userRepository.FindUserByEmail(ctx, email)
	if err != nil {
		return "", errors.ErrUserNotFound
	}
	// Compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.ErrInvalidPassword
	}
	_, token, err := config.TokenAuth.Encode(map[string]interface{}{"user_id": user.ID})
	if err != nil {
		return "", err
	}
	return token, nil
}

func (us *UserServiceImpl) ResetPassword(ctx context.Context, email string, password string) error {
	return nil
}
