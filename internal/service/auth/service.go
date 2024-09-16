package auth

import (
	"context"

	"github.com/lookandhate/course_auth/internal/service/model"
)

type Service struct {
}

func NewAuthService() *Service {
	return &Service{}
}

func (s *Service) Login(ctx context.Context, request model.LoginRequest) (string, error) {

	return "", nil
}
func (s *Service) GetRefreshToken(ctx context.Context, oldRefreshToken string) (string, error) {
	return "", nil
}
func (s *Service) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	return "", nil
}
