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
	userService "github.com/lookandhate/course_auth/internal/service/user"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	type userRepoMockFunc func(mc *minimock.Controller) repository.UserRepository
	type txManagerMockFunc func(f func(context.Context) error, mc *minimock.Controller) db.TxManager

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
		name               string
		args               args
		err                error
		userRepositoryMock userRepoMockFunc
		txManagerMock      txManagerMockFunc
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
				errTx = userRepoMock.DeleteUser(ctx, int(id))
				return errTx
			}, mc)
			serviceTest := userService.NewUserService(userRepoMock, txManagerMock)

			err := serviceTest.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)

		})
	}

}
