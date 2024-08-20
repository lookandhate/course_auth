package service

import (
	"context"
	"log"

	"github.com/lookandhate/course_auth/internal/service/model"
)

// Get validates user ID and after that tries to get user from repo.
func (s *Service) Get(ctx context.Context, id int) (*model.UserModel, error) {
	if err := s.validateID(id); err != nil {
		return nil, err
	}

	if err := s.checkUserExists(ctx, id); err != nil {
		return nil, err
	}

	user, err := s.cache.Get(ctx, id)
	if err == nil && user != nil {
		return user, nil
	}

	user, err = s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	// Do not think that we need to raise cache error above, just log it
	err = s.cache.Create(ctx, user)
	if err != nil {
		log.Default().Printf("Error when saving user to cache: %v", err)
	}

	return user, nil
}
