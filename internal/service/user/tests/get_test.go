package tests

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
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

		req = id

		expectedResponse = &model.UserModel{
			ID:        int(id),
			Name:      name,
			Email:     email,
			Role:      role,
			Password:  password,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
	)

	tests := []struct {
		name               string
		args               args
		expectedResult     *model.UserModel
		err                error
		userRepositoryMock userRepoMockFunc
		txManagerMock      txManagerMockFunc
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
			txManagerMock: func(f func(context.Context) error, mc *minimock.Controller) db.TxManager {
				mock := mocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Optional().Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})
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
				mock.GetUserMock.Optional().Expect(ctx, int(req)).Return(expectedResponse, nil)
				mock.CheckUserExistsMock.Expect(ctx, int(id)).Return(false, nil)
				return mock
			},
			txManagerMock: func(f func(context.Context) error, mc *minimock.Controller) db.TxManager {
				mock := mocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Optional().Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})
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
			serviceTest := userService.NewUserService(userRepoMock, txManagerMock)

			newID, err := serviceTest.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.expectedResult, newID)
		})
	}
}
