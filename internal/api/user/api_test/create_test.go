package api_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/lookandhate/course_auth/internal/api/user"
	"github.com/lookandhate/course_auth/internal/service"
	serviceMocks "github.com/lookandhate/course_auth/internal/service/mocks"
	"github.com/lookandhate/course_auth/internal/service/model"
	userApi "github.com/lookandhate/course_auth/pkg/auth_v1"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	var (
		ctx = context.Background()

		mc = minimock.NewController(t)

		email    = gofakeit.Email()
		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		password = gofakeit.Password(true, true, true, true, true, 10)
		role     = 1

		req = &userApi.CreateRequest{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            userApi.UserRole(role),
		}
		info = &model.CreateUserModel{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            model.UserRole(role),
		}
		res = &userApi.CreateResponse{
			Id: id,
		}
	)
	type args struct {
		ctx context.Context
		req *userApi.CreateRequest
	}

	tests := []struct {
		name            string
		args            args
		expectedResult  *userApi.CreateResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		// Test cases go below
		{
			name:           "success case",
			args:           args{ctx: ctx, req: req},
			expectedResult: res,
			err:            nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.RegisterMock.Expect(ctx, info).Return(id, nil)
				return mock
			},
		},
		{
			name:           "fail case",
			args:           args{ctx: ctx, req: req},
			expectedResult: nil,
			err:            service.ErrPasswordMismatch,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.RegisterMock.Expect(ctx, info).Return(id, service.ErrPasswordMismatch)
				return mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := user.NewAuthServer(userServiceMock)

			newID, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.expectedResult, newID)
		})
	}
}
