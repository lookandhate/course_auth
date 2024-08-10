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

func NewMockService(deps ...interface{}) *Service {
	service := Service{}
	for _, v := range deps {
		switch s := v.(type) {
		case repository.UserRepository:
			service.repo = s
		case db.TxManager:
			service.txManager = s
		}

	}
	return &service
}
