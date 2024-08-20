package cache

import (
	"context"

	model "github.com/lookandhate/course_auth/internal/service/model"
)

type UserCache interface {
	Create(ctx context.Context, user *model.UserModel) error
	Get(ctx context.Context, id int) (*model.UserModel, error)
	Delete(ctx context.Context, id int) error
}
