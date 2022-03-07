package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rodrigodev/jumia-go/src/internal/infrastructure"
	"github.com/rodrigodev/jumia-go/src/internal/phone/repository"
	"github.com/rodrigodev/jumia-go/src/internal/phone/service"
	"github.com/rodrigodev/jumia-go/src/internal/transport"

	_ "github.com/golang/mock/mockgen/model"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const (
	serviceName = "jumia-go"
)

func setup(ctx context.Context) (err error) {
	logger := infrastructure.Logger(ctx)

	db, err := infrastructure.GetDbConnection()
	if err != nil {
		logger.Fatal("Could not connect to database", zap.Error(err))
	}

	defer func() {
		if err != nil {
			logger.Error("startup", zap.Error(err))
		}
		_ = logger.Sync()
	}()

	phoneRepository, err := repository.NewPhoneRepository(db)
	if err != nil {
		logger.Fatal("error setting up the phone repository", zap.Error(err))
	}
	phoneService := service.NewPhoneService(phoneRepository)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		h, err := transport.New(
			transport.Health(),
			transport.Phone(phoneService),
			transport.Static(),
		)

		if err != nil {
			return err
		}

		return transport.ListenAndServe(ctx, ":8081", h)
	})

	logger.Info(fmt.Sprintf("starting %s at port %d", serviceName, 8081))

	logger.Info("shutdown", zap.Error(g.Wait()))
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(c)

	go func() {
		select {
		case <-c:
		case <-ctx.Done():
		}
		cancel()
	}()

	if err := setup(ctx); err != nil {
		os.Exit(1)
	}
}
