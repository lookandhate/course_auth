package service

import (
	"github.com/lookandhate/course_auth/internal/cache"
	"github.com/lookandhate/course_auth/internal/client"
	"github.com/lookandhate/course_platform_lib/pkg/db"

	"github.com/lookandhate/course_auth/internal/repository"
)

type Service struct {
	repo            repository.UserRepository
	cache           cache.UserCache
	txManager       db.TxManager
	passwordManager client.PasswordManager
}

// NewUserService creates Service with given repo.
func NewUserService(repo repository.UserRepository, manager db.TxManager, cache cache.UserCache, passwordManager client.PasswordManager) *Service {
	return &Service{
		repo:            repo,
		txManager:       manager,
		cache:           cache,
		passwordManager: passwordManager,
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
		case cache.UserCache:
			service.cache = s
		case client.PasswordManager:
			service.passwordManager = s
		}
	}

	return &service
}
