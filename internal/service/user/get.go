package service

import (
	"context"

	"github.com/lookandhate/course_auth/internal/service/model"
)

// GetUser validates user ID and after that tries to get user from repo.
func (s *Service) Get(ctx context.Context, id int) (*model.UserModel, error) {
	if err := s.validateID(id); err != nil {
		return nil, err
	}

	if err := s.checkUserExists(ctx, id); err != nil {
		return nil, err
	}
	return s.repo.GetUser(ctx, id)
}
