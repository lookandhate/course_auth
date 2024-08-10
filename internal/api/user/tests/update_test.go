package tests

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
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	var (
		ctx = context.Background()

		mc = minimock.NewController(t)

		email     = gofakeit.Email()
		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		password  = gofakeit.Password(true, true, true, true, true, 10)
		role      = 1
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		req = &userApi.UpdateRequest{
			Id:       id,
			Name:     &wrapperspb.StringValue{Value: name},
			Email:    &wrapperspb.StringValue{Value: email},
			Role:     userApi.UserRole(role),
			Password: &wrapperspb.StringValue{Value: password},
		}
		info = model.UpdateUserModel{
			ID:       int(id),
			Name:     &name,
			Email:    &email,
			Password: &password,
			Role:     role,
		}
		res             = &emptypb.Empty{}
		serviceResponse = &model.UserModel{
			Name:      name,
			Email:     email,
			Password:  password,
			Role:      role,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			ID:        int(id),
		}
	)

	type args struct {
		ctx context.Context
		req *userApi.UpdateRequest
	}

	tests := []struct {
		name            string
		args            args
		expectedResult  *emptypb.Empty
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
				mock.UpdateMock.Expect(ctx, &info).Return(serviceResponse, nil)
				return mock
			},
		},
		{
			name:           "fail case does not exist",
			args:           args{ctx: ctx, req: req},
			expectedResult: nil,
			err:            service.ErrUserDoesNotExist,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, &info).Return(nil, service.ErrUserDoesNotExist)
				return mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := user.NewAuthServer(userServiceMock)

			userData, err := api.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.expectedResult, userData)
		})
	}
}
