package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	userServer "github.com/lookandhate/course_auth/internal/api/user"
	"github.com/lookandhate/course_auth/internal/closer"
	"github.com/lookandhate/course_auth/internal/config"
	"github.com/lookandhate/course_auth/internal/repository"
	userRepo "github.com/lookandhate/course_auth/internal/repository/user"
	"github.com/lookandhate/course_auth/internal/service"
	userService "github.com/lookandhate/course_auth/internal/service/user"
)

// serviceProvider is a DI container.
type serviceProvider struct {
	appCfg *config.AppConfig
	pgPool *pgxpool.Pool

	userRepository repository.UserRepository

	userService    service.UserService
	userServerImpl *userServer.Server
}

// newServiceProvider creates plain serviceProvider.
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// AppCfg returns config.AppConfig.
func (s *serviceProvider) AppCfg() *config.AppConfig {
	if s.appCfg == nil {
		s.appCfg = config.MustLoad()
	}

	return s.appCfg
}

// UserRepository creates(if not exist) and returns repository.UserRepository instance.
func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepo.NewPostgresRepository(s.PgPool(ctx))
	}

	return s.userRepository
}

// UserService creates and returns service.UserService.
func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewUserService(s.UserRepository(ctx))
	}

	return s.userService
}

// PgPool returns and creates(if not exists) pgxpool.Pool.
func (s *serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.New(ctx, s.AppCfg().DB.GetDSN())
		if err != nil {
			log.Fatal(err)
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.pgPool = pool
	}

	err := s.pgPool.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return s.pgPool
}

func (s *serviceProvider) UserServerImpl(ctx context.Context) *userServer.Server {
	if s.userServerImpl == nil {
		s.userServerImpl = userServer.NewAuthServer(s.UserService(ctx))
	}
	return s.userServerImpl
}
