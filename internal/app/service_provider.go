package app

import (
	"context"
	"log"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	userServer "github.com/lookandhate/course_auth/internal/api/user"
	"github.com/lookandhate/course_auth/internal/cache"
	userCache "github.com/lookandhate/course_auth/internal/cache/user"
	"github.com/lookandhate/course_auth/internal/client"
	"github.com/lookandhate/course_auth/internal/client/crypto"
	"github.com/lookandhate/course_auth/internal/config"
	"github.com/lookandhate/course_auth/internal/repository"
	userRepo "github.com/lookandhate/course_auth/internal/repository/user"
	"github.com/lookandhate/course_auth/internal/service"
	userService "github.com/lookandhate/course_auth/internal/service/user"
	"github.com/lookandhate/course_platform_lib/pkg/closer"
	"github.com/lookandhate/course_platform_lib/pkg/db"
	"github.com/lookandhate/course_platform_lib/pkg/db/pg"
	"github.com/lookandhate/course_platform_lib/pkg/db/transaction"
)

// serviceProvider is a DI container.
type serviceProvider struct {
	appCfg *config.AppConfig

	dbClient           db.Client
	transactionManager db.TxManager

	redisPool *redigo.Pool

	userCache      cache.UserCache
	userRepository repository.UserRepository

	userService    service.UserService
	userServerImpl *userServer.Server

	passwordManager client.PasswordManager
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
		s.userService = userService.NewUserService(
			s.UserRepository(ctx), s.TxManager(ctx), s.UserCache(), s.PasswordManager(),
		)
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

func (s *serviceProvider) RedisPool() *redigo.Pool {
	if s.redisPool == nil {
		s.redisPool = &redigo.Pool{
			MaxIdle:     s.AppCfg().Redis.MaxIdle,
			IdleTimeout: time.Duration(s.AppCfg().Redis.IdleTimeout),
			Dial: func() (redigo.Conn, error) {
				return redigo.Dial("tcp", s.AppCfg().Redis.Address())
			},
		}
	}

	return s.redisPool
}

func (s *serviceProvider) UserCache() cache.UserCache {
	if s.userCache == nil {
		s.userCache = userCache.NewRedisCache(s.RedisPool(), s.AppCfg().Redis)
	}

	return s.userCache
}

func (s *serviceProvider) PasswordManager() client.PasswordManager {
	if s.passwordManager == nil {
		s.passwordManager = crypto.NewBCryptPasswordManager()
	}

	return s.passwordManager
}
