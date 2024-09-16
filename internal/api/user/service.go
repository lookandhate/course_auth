package user

import (
	"github.com/lookandhate/course_auth/internal/service"
	authAPI "github.com/lookandhate/course_auth/pkg/user_v1"
)

type Server struct {
	authAPI.UnimplementedAuthServer
	userService service.UserService
}

func NewAuthServer(service service.UserService) *Server {
	return &Server{
		userService: service,
	}
}
