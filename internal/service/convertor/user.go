package convertor

import (
	"time"

	cache "github.com/lookandhate/course_auth/internal/cache/model"
	repository "github.com/lookandhate/course_auth/internal/repository/model"
	service "github.com/lookandhate/course_auth/internal/service/model"
)

// ServiceCreateUserModelToRepoCreateUserModel converts model from service layer to repo layer.
func ServiceCreateUserModelToRepoCreateUserModel(user *service.CreateUserModel) *repository.CreateUserModel {
	if user == nil {
		return nil
	}
	return &repository.CreateUserModel{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     repository.UserRole(user.Role),
	}
}

// RepoUserModelToServiceUserModel converts model from repo layer to service layer.
func RepoUserModelToServiceUserModel(user *repository.UserModel) *service.UserModel {
	return &service.UserModel{
		ID:        user.ID,
		Name:      user.Name,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Role:      user.Role,
	}
}

// ServiceUserModelToCacheUserModel converts model from service layer to cache layer.
func ServiceUserModelToCacheUserModel(user *service.UserModel) *cache.UserModel {
	return &cache.UserModel{
		ID:          int64(user.ID),
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		Password:    user.Password,
		CreatedAtNS: user.CreatedAt.UnixNano(),
		UpdatedAtNS: user.UpdatedAt.UnixNano(),
	}
}

// CacheUserModelToServiceUserModel converts model from cache layer to service layer.
func CacheUserModelToServiceUserModel(user *cache.UserModel) *service.UserModel {
	return &service.UserModel{
		ID:        int(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: time.Unix(0, user.CreatedAtNS),
		UpdatedAt: time.Unix(0, user.UpdatedAtNS),
	}
}
