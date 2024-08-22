package service

import (
	"context"

	"github.com/lookandhate/course_auth/internal/service"
	"github.com/lookandhate/course_auth/internal/service/model"
)

// Register validates CreateUserModel, then passes it to repo layer and returns created user id.
func (s *Service) Register(ctx context.Context, user *model.CreateUserModel) (int, error) {
	if user == nil {
		return 0, service.ErrEmptyUser
	}
	// Check user role has been passed correctly
	if user.Role == model.UserUnknownRole {
		return 0, service.ErrInvalidRole
	}

	if user.PasswordConfirm != user.Password {
		return 0, service.ErrPasswordMismatch
	}

	var createdUserID int
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var err error

		user.Password, err = s.passwordManager.HashPassword(user.Password)
		if err != nil {
			return err
		}

		createdUserID, err = s.repo.CreateUser(ctx, user)
		if err != nil {
			return err
		}

		createdUser, err := s.repo.GetUser(ctx, createdUserID)
		if err != nil {
			return err
		}

		err = s.cache.Create(ctx, createdUser)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return createdUserID, nil
}
