package main

import (
	"context"
	"errors"
	"golang-url-shortener/internal/controller"
	"golang-url-shortener/internal/repository"
	"golang-url-shortener/internal/service"
	"golang-url-shortener/pkg/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	"go.uber.org/fx"
)

var ConnectionModule = fx.Module(
	"connection",
	fx.Provide(
		database.PostgresqlDatabaseProvider,
	),
)

var ServiceModule = fx.Module(
	"service",
	fx.Provide(
		service.NewShortService,
	),
)

var RepositoryModule = fx.Module(
	"repository",
	fx.Provide(
		repository.NewLinkRepo,
	),
)

func startHttp(lc fx.Lifecycle, ctrl *controller.Controller) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	engine.LoadHTMLGlob("public/*")
	engine.Static("/static", "./static")
	engine.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	engine.GET("/at/:link", ctrl.Find)
	engine.POST("/shorten", ctrl.Short)

	server := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Info().Ctx(ctx).Msg("Running API on port :8080...")
				err := server.ListenAndServe()
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error().Ctx(ctx).Err(err).Msg("Run app error")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			log.Info().Ctx(timeoutCtx).Msg("Shutting down server...")
			if err := server.Shutdown(timeoutCtx); err != nil {
				log.Error().Ctx(timeoutCtx).Err(err).Msg("Error shutting down server")
			} else {
				log.Info().Ctx(timeoutCtx).Msg("Server shutdown complete.")
			}
			return nil
		},
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg("No .env file")
	}
	
	appFx := fx.New(
		ConnectionModule,
		ServiceModule,
		RepositoryModule,
		fx.Provide(controller.NewController),
		fx.Invoke(startHttp),
	)
	appFx.Run()
}
