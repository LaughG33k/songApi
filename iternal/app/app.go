package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/LaughG33k/songApi/pkg"
	"github.com/gin-gonic/gin"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
)

type App struct {
	srvProvider *serviceProvider
	httpServe   *http.Server
	gin         *gin.Engine
}

func NewApp(ctx context.Context) (*App, error) {

	app := &App{}

	if err := app.initDeps(ctx); err != nil {
		return nil, err
	}

	return app, nil

}

func (a *App) Run() {

	pkg.Log.Info("начло работы")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := a.httpServe.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	if err := a.httpServe.Close(); err != nil {
		pkg.Log.Infof("закрытие сервера произошло с ошибкой: %s", err)
		return
	}

}

func (a *App) initDeps(ctx context.Context) error {

	deps := []func(context.Context) error{
		a.initServiceProvider,
		a.initMigrations,
		a.initHttpServer,
		a.initApi,
	}

	for _, fn := range deps {
		if err := fn(ctx); err != nil {
			return err
		}
	}

	return nil
}
func (a *App) initMigrations(_ context.Context) error {

	cfg := a.srvProvider.ConfigLoad()

	migrations, err := migrate.New(
		"file://iternal/migrations/song",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DB, "disable"),
	)

	if err != nil {
		return err
	}

	defer migrations.Close()

	if err := migrations.Up(); err != nil {

		if err.Error() == "no change" {
			return nil
		}

		if err := migrations.Drop(); err != nil {
			return err
		}
		return err
	}

	return nil
}
func (a *App) initServiceProvider(_ context.Context) error {

	a.srvProvider = newServiceProvider()

	return nil
}

func (a *App) initHttpServer(_ context.Context) error {

	cfg := a.srvProvider.ConfigLoad()

	a.httpServe = &http.Server{
		Addr:         cfg.Addr,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	router := gin.New()
	a.gin = router

	a.httpServe.Handler = router

	return nil

}

func (a *App) initApi(ctx context.Context) error {

	api := a.srvProvider.SongApi(ctx)

	a.gin.GET("/songs", api.GetAll)
	a.gin.GET("/text", api.GetText)
	a.gin.POST("/add", api.Add)
	a.gin.DELETE("/delete/:group/:song", api.Delete)
	a.gin.PATCH("/edit/:group/:song", api.Edit)

	return nil
}
