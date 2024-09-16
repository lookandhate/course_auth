package auth

import (
	"github.com/lookandhate/course_auth/internal/service"
	authAPI "github.com/lookandhate/course_auth/pkg/auth_v1"
)

type Server struct {
	authAPI.UnimplementedAuthServer
	authService service.AuthService
}

func NewAuthServer(authService service.AuthService) *Server {
	return &Server{
		authService: authService,
	}
}
