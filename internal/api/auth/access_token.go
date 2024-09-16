package auth

import (
	"context"

	"github.com/lookandhate/course_auth/pkg/auth_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Server) GetAccessToken(context.Context, *auth_v1.GetAccessTokenRequest) (*auth_v1.GetAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
