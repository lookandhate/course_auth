package service

import (
	"context"
	"errors"
	"log"

	"github.com/lookandhate/course_auth/internal/service"
	"github.com/lookandhate/course_auth/internal/service/model"
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

	var updatedUser *model.UserModel
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var err error
		updatedUser, err = s.repo.UpdateUser(ctx, user)
		return err
	})
	if err != nil {
		return nil, err
	}

	// Invalidate user cache
	err = s.cache.Delete(ctx, updatedUser.ID)
	if err != nil {
		log.Printf("Error when deleting user from cache: %v", err)
	}

	return updatedUser, nil
}
