package convertor

import (
	repository "github.com/lookandhate/course_auth/internal/repository/model"
	"github.com/lookandhate/course_auth/internal/service/model"
)

func UserRepoToService(user *repository.UserModel) *model.UserModel {
	return &model.UserModel{
		ID:        user.ID,
		Name:      user.Name,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

}
