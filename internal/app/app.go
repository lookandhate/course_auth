package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/lookandhate/course_auth/pkg/auth_v1"
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

	ctx, cancel := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(1)

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

	gracefulShutdown(ctx, cancel, wg)

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

	auth_v1.RegisterAuthServer(a.grpcServer, a.serviceProvider.UserServerImpl(ctx))

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

func gracefulShutdown(ctx context.Context, cancelFunc context.CancelFunc, wg *sync.WaitGroup) {
	select {
	case <-ctx.Done():
		log.Printf("terminating due to context cancelling")
		break
	case <-waitSignal():
		log.Printf("terminating via term sig")
		break

	}
	cancelFunc()

	if wg != nil {
		wg.Wait()
		fmt.Printf("After wait")
	}
}

func waitSignal() chan os.Signal {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	return sigterm
}
