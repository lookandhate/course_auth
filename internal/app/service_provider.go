package app

import (
	"context"
	"log"

	userServer "github.com/lookandhate/course_auth/internal/api/user"
	"github.com/lookandhate/course_auth/internal/client/db"
	"github.com/lookandhate/course_auth/internal/client/db/pg"
	"github.com/lookandhate/course_auth/internal/client/transaction"
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

	dbClient           db.Client
	transactionManager db.TxManager

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
		s.userRepository = userRepo.NewPostgresRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

// UserService creates and returns service.UserService.
func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewUserService(s.UserRepository(ctx), s.TxManager(ctx))
	}

	return s.userService
}

func (s *serviceProvider) UserServerImpl(ctx context.Context) *userServer.Server {
	if s.userServerImpl == nil {
		s.userServerImpl = userServer.NewAuthServer(s.UserService(ctx))
	}
	return s.userServerImpl
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.AppCfg().DB.GetDSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.transactionManager == nil {
		s.transactionManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.transactionManager
}
