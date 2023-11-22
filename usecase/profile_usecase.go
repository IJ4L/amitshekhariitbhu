package usecase

import (
	"context"
	"time"

	"architecture.com/domain"
)

type profileUseCase struct {
	UserRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(ur domain.UserRepository, timeout time.Duration) domain.ProfileUsecase {
	return &profileUseCase{
		UserRepository: ur,
		contextTimeout: timeout,
	}
}

// GetProfileByID implements domain.ProfileUsecase.
func (pu *profileUseCase) GetProfileByID(c context.Context, userID string) (*domain.Profile, error) {
	_, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	user, err := pu.UserRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return &domain.Profile{Name: user.Name, Email: user.Email}, nil
}
