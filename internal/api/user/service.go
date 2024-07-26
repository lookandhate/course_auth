package user

import (
	"github.com/lookandhate/course_auth/internal/service"
	"github.com/lookandhate/course_auth/pkg/auth_v1"
)

// grpc server Implementation.
type Implementation struct {
	auth_v1.UnimplementedAuthServer
	service service.UserService
}

// NewImplementation creates grpc server Implementation.
func NewImplementation(service service.UserService) *Implementation {
	return &Implementation{service: service}
}
