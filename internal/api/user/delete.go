package user

import (
	"context"

	authAPI "github.com/lookandhate/course_auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Delete(ctx context.Context, request *authAPI.DeleteRequest) (*emptypb.Empty, error) {
	err := s.userService.Delete(ctx, int(request.GetId()))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
