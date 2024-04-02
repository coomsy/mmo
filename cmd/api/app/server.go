package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coomsy/mmo/internal/config"
	"github.com/coomsy/mmo/internal/http/routes"
	"github.com/coomsy/mmo/pkg/logger"
)

type App struct {
	HttpServer *http.Server
}

func NewApp() (*App, error) {
	// db

	// handler
	router := http.NewServeMux()

	router.HandleFunc("/", routes.Root)
	routes.

	// API Routes
	//api := router.Group("api")
	//api.GET("/", routes.RootHandler)
	//routes.NewUsersRoute(api, conn, jwtService, redisCache, ristrettoCache, authMiddleware, mailerService).Routes()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 20 MB, KC cookies w/o redis prolly gonna break this
	}

	return &App{
		HttpServer: server,
	}, nil
}

func (a *App) Run() (err error) {
	// graceful shutdown
	go func() {
		logger.Info(fmt.Sprintf("Success to listen and serve on :%d", config.AppConfig.Port))
		if err := a.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(fmt.Sprintf("Failed to listen and serve: %+v", err))
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	logger.Info("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := a.HttpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("error when shutdown server: %v", err)
	}

	<-ctx.Done()
	logger.Info("Server exiting")
	return
}
