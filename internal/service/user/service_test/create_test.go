package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/lookandhate/course_auth/internal/cache"
	cacheMocks "github.com/lookandhate/course_auth/internal/cache/mocks"
	"github.com/lookandhate/course_auth/internal/client"
	clientMocks "github.com/lookandhate/course_auth/internal/client/mocks"
	"github.com/lookandhate/course_auth/internal/repository"
	repoMocks "github.com/lookandhate/course_auth/internal/repository/mocks"
	"github.com/lookandhate/course_auth/internal/service"
	"github.com/lookandhate/course_auth/internal/service/model"
	userService "github.com/lookandhate/course_auth/internal/service/user"
	"github.com/lookandhate/course_platform_lib/pkg/db"
	"github.com/lookandhate/course_platform_lib/pkg/db/mocks"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	type userRepoMockFunc func(mc *minimock.Controller) repository.UserRepository
	type txManagerMockFunc func(f func(context.Context) error, mc *minimock.Controller) db.TxManager
	type userCacheMockFunc func(mc *minimock.Controller) cache.UserCache
	type passwordManagerMockFunc func(mc *minimock.Controller) client.PasswordManager
	type args struct {
		ctx context.Context
		req *model.CreateUserModel
	}

	t.Parallel()

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, false, gofakeit.Number(5, 15))
		role     = gofakeit.Number(1, 2)
		timeNow  = time.Now()

		reqSuccess = &model.CreateUserModel{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            model.UserRole(role),
		}
		responseSuccess = &model.UserModel{
			ID:        int(id),
			Name:      name,
			Email:     email,
			Role:      role,
			Password:  password,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		}
		reqPassMismatch = &model.CreateUserModel{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: gofakeit.Name(),
			Role:            model.UserRole(role),
		}
	)
	tests := []struct {
		name                string
		args                args
		expectedResult      int
		err                 error
		userRepositoryMock  userRepoMockFunc
		txManagerMock       txManagerMockFunc
		userCacheMock       userCacheMockFunc
		passwordManagerMock passwordManagerMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: reqSuccess,
			},
			expectedResult: int(id),
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateUserMock.Expect(ctx, reqSuccess).Return(int(id), nil)
				mock.GetUserMock.Expect(ctx, int(id)).Return(responseSuccess, nil)
				return mock
			},
			txManagerMock: func(_ func(context.Context) error, mc *minimock.Controller) db.TxManager {
				mock := mocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) error {
					return f(ctx)
				})
				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.CreateMock.Expect(ctx, responseSuccess).Return(nil)
				return mock
			},
			passwordManagerMock: func(mc *minimock.Controller) client.PasswordManager {
				mock := clientMocks.NewPasswordManagerMock(mc)
				mock.HashPasswordMock.Expect(password).Return(password, nil)
				return mock
			},
		},
		{
			name: "fail with password",
			args: args{
				ctx: ctx,
				req: reqPassMismatch,
			},
			expectedResult: 0,
			err:            service.ErrPasswordMismatch,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateUserMock.Optional().
					Expect(ctx, reqPassMismatch).Return(0, nil)
				return mock
			},
			txManagerMock: func(_ func(context.Context) error, mc *minimock.Controller) db.TxManager {
				mock := mocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Optional().Set(func(ctx context.Context, f db.Handler) error {
					return f(ctx)
				})
				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.CreateMock.Optional().Return(nil)
				return mock
			},
			passwordManagerMock: func(mc *minimock.Controller) client.PasswordManager {
				mock := clientMocks.NewPasswordManagerMock(mc)
				mock.HashPasswordMock.Optional().Expect(password).Return(password, nil)
				return mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepoMock := tt.userRepositoryMock(mc)
			txManagerMock := tt.txManagerMock(func(ctx context.Context) error {
				var errTx error
				idInt, errTx := userRepoMock.CreateUser(ctx, tt.args.req)
				id = int64(idInt)
				if errTx != nil {
					return errTx
				}
				return nil
			}, mc)
			userCacheMock := tt.userCacheMock(mc)
			passwordManagerMock := tt.passwordManagerMock(mc)

			serviceTest := userService.NewUserService(userRepoMock, txManagerMock, userCacheMock, passwordManagerMock)

			newID, err := serviceTest.Register(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.expectedResult, newID)
		})
	}
}
