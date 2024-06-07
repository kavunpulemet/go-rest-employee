package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go-rest-employee/config"
	"go-rest-employee/pkg/api"
	"go-rest-employee/pkg/repository"
	"go-rest-employee/pkg/service"
	"go.uber.org/zap"
)

type App struct {
	ctx        context.Context
	server     *api.Server
	repository *sqlx.DB
	logger     *zap.SugaredLogger
	settings   config.Settings
}

func NewApp(ctx context.Context, logger *zap.SugaredLogger, settings config.Settings) *App {
	return &App{
		ctx:      ctx,
		logger:   logger,
		settings: settings,
	}
}

func (a *App) InitDatabase() error {
	var err error
	a.repository, err = sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		a.settings.Host, a.settings.Port, a.settings.Username, a.settings.DBName, a.settings.Password, a.settings.SSLMode))
	if err != nil {
		return err
	}

	err = a.repository.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) InitService() {
	s := service.NewEmployeeService(repository.NewRepository(a.repository))
	a.server = api.NewServer()
	a.server.HandleEmployees(s)
}

func (a *App) Run() error {
	go func() {
		if err := a.server.Run(); err != nil {
			a.logger.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	a.logger.Info("run server")
	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	err := a.server.Shutdown(ctx)
	if err != nil {
		a.logger.Errorf("Failed to disconnect from server %v", err)
		return err
	}

	err = a.repository.Close()
	if err != nil {
		a.logger.Errorf("failed to disconnect from bd %v", err)
	}

	a.logger.Info("server shut down successfully")
	return nil
}
