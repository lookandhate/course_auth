package app

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/lookandhate/course_auth/pkg/user_v1"
	"github.com/lookandhate/course_platform_lib/pkg/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2) //nolint:mnd // not really a magic number

	go func() {
		defer wg.Done()

		err := a.serviceProvider.UserSaverConsumer(ctx).RunConsumer(ctx)
		if err != nil {
			log.Printf("error runing UserSaverConsumer: %v\n", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runGRPCServer()
		if err != nil {
			log.Fatalf("error runing grpc app: %v\n", err)
		}
	}()

	wg.Wait()

	return nil
}

// initDeps initialize dependencies.
func (a *App) initDeps(ctx context.Context) error {
	initFuncs := []func(context.Context) error{

		a.initServiceProvider, a.initGRPCServer,
	}
	for _, f := range initFuncs {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	user_v1.RegisterAuthServer(a.grpcServer, a.serviceProvider.UserServerImpl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	serveAddress := a.serviceProvider.AppCfg().GPRC.Address()
	log.Printf("GRPC server is running on %s", serveAddress)

	listener, err := net.Listen("tcp", serveAddress)
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}
