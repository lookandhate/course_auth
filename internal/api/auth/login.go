package auth

import (
	"context"

	"github.com/lookandhate/course_auth/pkg/auth_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Server) Login(context.Context, *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
