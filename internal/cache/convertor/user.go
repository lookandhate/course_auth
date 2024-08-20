package convertor

import (
	"time"

	cache "github.com/lookandhate/course_auth/internal/cache/model"
	service "github.com/lookandhate/course_auth/internal/service/model"
)

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
