package app

import (
	"calc/internal/api/conversion"
	"calc/internal/config"
	"calc/internal/database"
	// swagger docs
	_ "calc/internal/docs"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/jmoiron/sqlx"
	"time"
)

// App application
type App struct {
	serviceProvider *serviceProvider
	fiberApp        *fiber.App
	db              *sqlx.DB
	cfg             *config.Config
}

// NewApp create new application
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run application
func (a *App) Run(ctx context.Context) error {
	go a.serviceProvider.conversionService.Run(ctx)
	return a.runRestServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initDB,
		a.initServiceProvider,
		a.initRestServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	a.cfg = cfg
	return nil
}

func (a *App) initDB(ctx context.Context) error {
	var err error
	a.db, err = database.Connect(ctx)
	if err != nil {
		return err
	}

	err = database.Migrate(a.db)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.db, a.cfg)
	return nil
}

// @title Calculator
// @version 1.0
// @description Calculator.

// @host localhost:8080
// @BasePath /
// @schemes http
func (a *App) initRestServer(_ context.Context) error {
	a.fiberApp = fiber.New(fiber.Config{
		AppName:      "Calculator",
		ServerHeader: "Fiber",
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	})

	a.fiberApp.Group("/swagger").Get("*", swagger.HandlerDefault)
	a.fiberApp.Group("/api/v1").Get("/convert", conversion.NewHandler(a.serviceProvider.ConvertService()))

	return nil
}

func (a *App) runRestServer() error {
	if err := a.fiberApp.Listen(":8080"); err != nil {
		return err
	}
	return nil
}
