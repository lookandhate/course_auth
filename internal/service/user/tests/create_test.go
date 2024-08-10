package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/lookandhate/course_auth/internal/convertor"
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

		req = &model.CreateUserModel{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            model.UserRole(role),
		}
		reqPassMissmath = &model.CreateUserModel{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: gofakeit.Name(),
			Role:            model.UserRole(role),
		}
	)
	tests := []struct {
		name               string
		args               args
		expectedResult     int
		err                error
		userRepositoryMock userRepoMockFunc
		txManagerMock      txManagerMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: req,
			},
			expectedResult: int(id),
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateUserMock.Expect(ctx, convertor.CreateUserModelToRepo(req)).Return(int(id), nil)
				return mock
			},
			txManagerMock: func(f func(context.Context) error, mc *minimock.Controller) db.TxManager {
				mock := mocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})
				return mock
			},
		},
		{
			name: "fail with password",
			args: args{
				ctx: ctx,
				req: reqPassMissmath,
			},
			expectedResult: 0,
			err:            service.ErrPasswordMismatch,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateUserMock.Optional().Expect(ctx, convertor.CreateUserModelToRepo(reqPassMissmath)).Return(0, nil)
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
				idInt, errTx := userRepoMock.CreateUser(ctx, convertor.CreateUserModelToRepo(req))
				id = int64(idInt)
				if errTx != nil {
					return errTx
				}
				return nil
			}, mc)
			serviceTest := userService.NewUserService(userRepoMock, txManagerMock)

			newID, err := serviceTest.Register(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.expectedResult, newID)
		})
	}

}
