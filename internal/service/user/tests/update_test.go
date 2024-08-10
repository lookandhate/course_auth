package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/lookandhate/course_auth/internal/client/db"
	"github.com/lookandhate/course_auth/internal/client/db/mocks"
	"github.com/lookandhate/course_auth/internal/repository"
	repoMocks "github.com/lookandhate/course_auth/internal/repository/mocks"
	"github.com/lookandhate/course_auth/internal/service/model"
	userService "github.com/lookandhate/course_auth/internal/service/user"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	type userRepoMockFunc func(mc *minimock.Controller) repository.UserRepository
	type txManagerMockFunc func(f func(context.Context) error, mc *minimock.Controller) db.TxManager
	type args struct {
		ctx context.Context
		req *model.UpdateUserModel
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

		req = &model.UpdateUserModel{
			ID:       int(id),
			Name:     &name,
			Email:    &email,
			Password: &password,
			Role:     role,
		}

		expectedResponse = &model.UserModel{
			Name:     name,
			Email:    email,
			Role:     role,
			Password: password,
			ID:       int(id),
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
				ctx: ctx,
				req: req,
			},
			expectedResult: expectedResponse,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.UpdateUserMock.Expect(ctx, req).Return(expectedResponse, nil)
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepoMock := tt.userRepositoryMock(mc)
			txManagerMock := tt.txManagerMock(func(ctx context.Context) error {
				var errTx error
				_, errTx = userRepoMock.UpdateUser(ctx, req)
				if errTx != nil {
					return errTx
				}
				return nil
			}, mc)
			serviceTest := userService.NewUserService(userRepoMock, txManagerMock)

			newID, err := serviceTest.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.expectedResult, newID)
		})
	}
}
