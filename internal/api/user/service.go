package user

import (
	"github.com/lookandhate/course_auth/internal/service"
	userAPI "github.com/lookandhate/course_auth/pkg/user_v1"
)

type Server struct {
	userAPI.UnimplementedUserServer
	userService service.UserService
}

func NewUserServer(service service.UserService) *Server {
	return &Server{
		userService: service,
	}
}
