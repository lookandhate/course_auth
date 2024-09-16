package auth

import (
	"context"

	"github.com/lookandhate/course_auth/pkg/auth_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Server) GetRefreshToken(context.Context, *auth_v1.GetRefreshTokenRequest) (*auth_v1.GetRefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRefreshToken not implemented")
}
