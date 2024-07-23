package service

import "github.com/lookandhate/microservice-courese/auth/internal/repository"

type Service struct {
	repo repository.UserRepository
}

// NewUserService creates Service with given repo.
func NewUserService(repo repository.UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}
