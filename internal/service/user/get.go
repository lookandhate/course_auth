package service

import (
	"context"
	"log"

	"github.com/lookandhate/course_auth/internal/service/convertor"
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

	userFromCache, err := s.cache.Get(ctx, id)
	if err == nil && userFromCache != nil {
		return convertor.CacheUserModelToServiceUserModel(userFromCache), nil
	}

	userFromRepo, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	userService := convertor.RepoUserModelToServiceUserModel(userFromRepo)

	// Do not think that we need to raise cache error above, just log it
	err = s.cache.Create(ctx, convertor.ServiceUserModelToCacheUserModel(userService))
	if err != nil {
		log.Default().Printf("Error when saving user to cache: %v", err)
	}

	return userService, nil
}
