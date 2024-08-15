package api_tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/lookandhate/course_auth/internal/api/user"
	"github.com/lookandhate/course_auth/internal/service"
	serviceMocks "github.com/lookandhate/course_auth/internal/service/mocks"
	userApi "github.com/lookandhate/course_auth/pkg/auth_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	var (
		ctx = context.Background()

		mc = minimock.NewController(t)

		id  = gofakeit.Int64()
		req = &userApi.DeleteRequest{
			Id: id,
		}
		res = &emptypb.Empty{}
	)
	type args struct {
		ctx context.Context
		req *userApi.DeleteRequest
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
				mock.DeleteMock.Expect(ctx, int(id)).Return(nil)
				return mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := user.NewAuthServer(userServiceMock)

			response, err := api.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.expectedResult, response)
		})
	}
}
