package user

import (
	"context"

	"github.com/lookandhate/course_auth/internal/api/convertor"
	authAPI "github.com/lookandhate/course_auth/pkg/auth_v1"
)

func (s *Server) Create(ctx context.Context, request *authAPI.CreateRequest) (*authAPI.CreateResponse, error) {
	userID, err := s.userService.Register(ctx, convertor.CreateUserFromProto(request))
	if err != nil {
		return nil, err
	}

	return &authAPI.CreateResponse{Id: int64(userID)}, nil
}
