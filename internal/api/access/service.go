package access

import (
	"github.com/lookandhate/course_auth/internal/service"
	accessAPI "github.com/lookandhate/course_auth/pkg/access_v1"
)

type Server struct {
	accessAPI.UnimplementedAccessServer
	accessService service.AccessService
}

func NewAccessServer(accessService service.AccessService) *Server {
	return &Server{
		accessService: accessService,
	}
}
