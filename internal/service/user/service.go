package service

import (
	"github.com/lookandhate/course_auth/internal/client/db"
	"github.com/lookandhate/course_auth/internal/repository"
)

type Service struct {
	repo      repository.UserRepository
	txManager db.TxManager
}

// NewUserService creates Service with given repo.
func NewUserService(repo repository.UserRepository, manager db.TxManager) *Service {
	return &Service{
		repo:      repo,
		txManager: manager,
	}
}
