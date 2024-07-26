package service

import (
	"context"

	"github.com/lookandhate/course_auth/internal/convertor"
	"github.com/lookandhate/course_auth/internal/service"
	"github.com/lookandhate/course_auth/internal/service/model"
)

// RegisterUser validates CreateUserModel, then passes it to repo layer and returns created user id.
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

	createUserID, err := s.repo.CreateUser(ctx, convertor.CreateUserModelToRepo(user))
	if err != nil {
		return 0, err
	}

	return createUserID, nil
}
