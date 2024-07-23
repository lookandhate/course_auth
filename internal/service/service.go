package service

import (
	"context"

	"github.com/lookandhate/microservice-courese/auth/internal/service/model"
)

type UserService interface {
	Register(ctx context.Context, user *model.CreateUserModel) (id int, err error)
	Get(ctx context.Context, id int) (user *model.UserModel, err error)
	Update(ctx context.Context, user *model.UpdateUserModel) (updatedUser *model.UserModel, err error)
	Delete(ctx context.Context, id int) (err error)
}
