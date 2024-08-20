package user

import (
	"context"

	"github.com/lookandhate/course_auth/internal/service/convertor"
	authAPI "github.com/lookandhate/course_auth/pkg/auth_v1"
)

func (s *Server) Get(ctx context.Context, request *authAPI.GetRequest) (*authAPI.GetResponse, error) {
	user, err := s.userService.Get(ctx, int(request.GetId()))
	if err != nil {
		return nil, err
	}

	return convertor.UserModelToGetResponseProto(user), err
}
