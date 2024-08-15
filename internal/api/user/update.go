package user

import (
	"context"

	"github.com/lookandhate/course_auth/internal/api/convertor"
	authAPI "github.com/lookandhate/course_auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Update(ctx context.Context, request *authAPI.UpdateRequest) (*emptypb.Empty, error) {
	_, err := s.userService.Update(ctx, convertor.UserUpdateFromProto(request))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
