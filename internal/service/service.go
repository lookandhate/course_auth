package service

import (
	"context"

	"github.com/lookandhate/course_auth/internal/service/model"
)

type UserService interface {
	Register(ctx context.Context, user *model.CreateUserModel) (id int, err error)
	Get(ctx context.Context, id int) (user *model.UserModel, err error)
	Update(ctx context.Context, user *model.UpdateUserModel) (updatedUser *model.UserModel, err error)
	Delete(ctx context.Context, id int) (err error)
}

type AuthService interface {
	Login(ctx context.Context, request model.LoginRequest) (refreshToken string, err error)
	GetRefreshToken(ctx context.Context, oldRefreshToken string) (refreshToken string, err error)
	GetAccessToken(ctx context.Context, refreshToken string) (accessToken string, err error)
}

type AccessService interface {
}
