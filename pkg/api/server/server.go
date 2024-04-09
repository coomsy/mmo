package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	v1User "github.com/coomsy/mmo/pkg/api/handlers/user/v1"
	"github.com/coomsy/mmo/pkg/config"
	"github.com/coomsy/mmo/pkg/log"
)

type HTTPServer struct {
	httpServer *http.Server
	mux        *chi.Mux
}

func NewServer() (*HTTPServer, error) {
	// https://github.com/grafana/grafana/blob/main/pkg/api/api.go

	r := newRouter()

	SetupRoutes(r)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler: r,
		// avoid Slowloris attacks.
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 20 MB, KC cookies w/o redis prolly gonna break this
	}

	return &HTTPServer{
		httpServer: server,
		mux:        r,
	}, nil
}

func (hs *HTTPServer) Run() (err error) {
	// graceful shutdown
	go func() {
		log.Infof("Success to listen and serve on :%d", config.AppConfig.Port)

		if err := hs.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	log.Info("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := hs.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("error when shutdown server: %v", err)
	}

	<-ctx.Done()
	log.Info("Server exiting")
	return
}

func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	return r
}

func SetupRoutes(mux *chi.Mux) {

	// Need way of passing db conn to handler,
	//	configure handler struct somehow before hand?

	// api, dbconn, redisCache, ristrettoCache, authMiddleware, mailerService

	userHandler := v1User.SetupHandler()

	mux.Route("/v1", func(r chi.Router) {
		mux.Get("/user", userHandler.GetUser)
		mux.Get("/user/{userId}", userHandler.GetUserById)
	})
}
