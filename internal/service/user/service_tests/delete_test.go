package service_tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/lookandhate/course_auth/internal/cache"
	cacheMocks "github.com/lookandhate/course_auth/internal/cache/mocks"
	"github.com/lookandhate/course_auth/internal/client"
	clientMocks "github.com/lookandhate/course_auth/internal/client/mocks"
	"github.com/lookandhate/course_auth/internal/repository"
	repoMocks "github.com/lookandhate/course_auth/internal/repository/mocks"
	userService "github.com/lookandhate/course_auth/internal/service/user"
	"github.com/lookandhate/course_platform_lib/pkg/db"
	"github.com/lookandhate/course_platform_lib/pkg/db/mocks"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
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
		id  = gofakeit.Uint32()
		ctx = context.Background()
		mc  = minimock.NewController(t)
	)
	tests := []struct {
		name                string
		args                args
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
				req: int(id),
			},
			err: nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.DeleteUserMock.Expect(ctx, int(id)).Return(nil)
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
				mock := cacheMocks.NewUserCacheMock(t)
				mock.DeleteMock.Expect(ctx, int(id)).Return(nil)
				return mock
			},
			passwordManagerMock: func(mc *minimock.Controller) client.PasswordManager {
				mock := clientMocks.NewPasswordManagerMock(t)
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
				return userRepoMock.DeleteUser(ctx, int(id))
			}, mc)
			userCacheMock := tt.userCacheMock(mc)
			passwordManagerMock := tt.passwordManagerMock(mc)

			serviceTest := userService.NewUserService(userRepoMock, txManagerMock, userCacheMock, passwordManagerMock)

			err := serviceTest.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
