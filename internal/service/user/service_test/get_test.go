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

func TestGet(t *testing.T) {
	type userRepoMockFunc func(mc *minimock.Controller) repository.UserRepository
	type txManagerMockFunc func(f func(context.Context) error, mc *minimock.Controller) db.TxManager
	type userCacheMockFunc func(mc *minimock.Controller) cache.UserCache
	type passwordManagerMockFunc func(mc *minimock.Controller) client.PasswordManager

	type args struct {
		ctx context.Context
		req int
	}

	t.Parallel()
	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Uint32()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, false, gofakeit.Number(5, 15))
		role     = gofakeit.Number(1, 2)
		timeNow  = time.Now()

		req = id

		expectedResponse = &model.UserModel{
			ID:        int(id),
			Name:      name,
			Email:     email,
			Role:      role,
			Password:  password,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		}
	)

	tests := []struct {
		name                string
		args                args
		expectedResult      *model.UserModel
		err                 error
		userRepositoryMock  userRepoMockFunc
		txManagerMock       txManagerMockFunc
		userCacheMock       userCacheMockFunc
		passwordManagerMock passwordManagerMockFunc
	}{
		{
			name: "success",
			args: args{
				req: int(id),
				ctx: ctx,
			},
			expectedResult: expectedResponse,
			err:            nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetUserMock.Expect(ctx, int(req)).Return(expectedResponse, nil)
				mock.CheckUserExistsMock.Expect(ctx, int(id)).Return(true, nil)
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
				mock.GetMock.Optional().Expect(ctx, int(req)).Return(nil, nil)
				mock.CreateMock.Optional().Expect(ctx, expectedResponse).Return(nil)
				return mock
			},
			passwordManagerMock: func(mc *minimock.Controller) client.PasswordManager {
				mock := clientMocks.NewPasswordManagerMock(mc)
				mock.HashPasswordMock.Optional()
				return mock
			},
		},
		{
			name: "error user does not exist",
			args: args{
				req: int(id),
				ctx: ctx,
			},
			expectedResult: nil,
			err:            service.ErrUserDoesNotExist,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetUserMock.Optional().Expect(ctx, int(req)).Return(nil, nil)
				mock.CheckUserExistsMock.Expect(ctx, int(id)).Return(false, nil)
				return mock
			},
			userCacheMock: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.GetMock.Optional().Expect(ctx, int(req)).Return(nil, nil)
				return mock
			},
			txManagerMock: func(_ func(context.Context) error, mc *minimock.Controller) db.TxManager {
				mock := mocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Optional().Set(func(ctx context.Context, f db.Handler) error {
					return f(ctx)
				})
				return mock
			},
			passwordManagerMock: func(mc *minimock.Controller) client.PasswordManager {
				mock := clientMocks.NewPasswordManagerMock(mc)
				mock.HashPasswordMock.Optional()
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
				_, errTx = userRepoMock.GetUser(ctx, int(req))
				if errTx != nil {
					return errTx
				}
				return nil
			}, mc)
			userCacheMock := tt.userCacheMock(mc)
			passwordManagerMock := tt.passwordManagerMock(mc)

			serviceTest := userService.NewUserService(userRepoMock, txManagerMock, userCacheMock, passwordManagerMock)

			newID, err := serviceTest.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.expectedResult, newID)
		})
	}
}
