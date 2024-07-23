package service

import (
	"context"
	"errors"

	"github.com/lookandhate/microservice-courese/auth/internal/service"
	"github.com/lookandhate/microservice-courese/auth/internal/service/model"
)

// UpdateUser validates passed user data and updates user info.
func (s *Service) Update(ctx context.Context, user *model.UpdateUserModel) (*model.UserModel, error) {
	if user == nil {
		err := errors.New("user is nil")
		return nil, err
	}

	if user.Role == int(model.UserUnknownRole) {
		err := service.ErrInvalidRole
		return nil, err
	}

	return s.repo.UpdateUser(ctx, user)
}
