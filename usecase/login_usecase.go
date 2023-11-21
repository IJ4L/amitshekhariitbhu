package usecase

import (
	"context"
	"time"

	"architecture.com/domain"
	"architecture.com/internal/tokenutil"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	ContextTimeout time.Duration
}

func NewLoginUsecase(ur domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: ur,
		ContextTimeout: timeout,
	}
}

// GetUserByEmail implements domain.LoginUsecase.
func (ur *loginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	_, cancel := context.WithTimeout(c, time.Second)
	defer cancel()
	return ur.userRepository.GetByEmail(email)
}

// CreateAccessToken implements domain.LoginUsecase.
func (ur *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

// CreateRefreshToken implements domain.LoginUsecase.
func (ur *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
