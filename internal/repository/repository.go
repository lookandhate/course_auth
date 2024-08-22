package repository

import (
	"context"

	"github.com/lookandhate/course_auth/internal/service/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.CreateUserModel) (int, error)
	GetUser(ctx context.Context, id int) (*model.UserModel, error)
	UpdateUser(ctx context.Context, updateUser *model.UpdateUserModel) (*model.UserModel, error)
	DeleteUser(ctx context.Context, id int) error
	CheckUserExists(ctx context.Context, id int) (bool, error)
}
