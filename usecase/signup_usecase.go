package usecase

import (
	"context"
	"time"

	"architecture.com/domain"
	"architecture.com/internal/tokenutil"
)

type signUpUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignUpUsecase(ur domain.UserRepository, timeout time.Duration) domain.SignupUsecase {
	return &signUpUsecase{
		userRepository: ur,
		contextTimeout: timeout,
	}
}

// Create implements domain.SignupUsecase.
func (su *signUpUsecase) Create(c context.Context, user *domain.User) error {
	_, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(user)
}

// CreateAccessToken implements domain.SignupUsecase.
func (su *signUpUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

// CreateRefreshToken implements domain.SignupUsecase.
func (su *signUpUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

// GetUserByEmail implements domain.SignupUsecase.
func (su *signUpUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	_, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(email)
}