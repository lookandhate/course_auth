package repository

import (
	"context"

	repository "github.com/lookandhate/course_auth/internal/repository/model"
	"github.com/lookandhate/course_auth/internal/service/model"
)

type UserRepository interface {
	CreateUser(context context.Context, user *repository.CreateUserModel) (int, error)
	GetUser(context context.Context, id int) (*model.UserModel, error)
	UpdateUser(context context.Context, updateUser *model.UpdateUserModel) (*model.UserModel, error)
	DeleteUser(context context.Context, id int) error
	CheckUserExists(context context.Context, id int) (bool, error)
}
