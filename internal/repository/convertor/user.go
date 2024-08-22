package convertor

import (
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
