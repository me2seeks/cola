package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/me2seeks/cola/config"
	"github.com/me2seeks/cola/internal/pkg/logger"

	"github.com/me2seeks/cola/internal/router"
)

func main() {
	server := &http.Server{
		Addr:    config.Cfg.Host + ":" + strconv.Itoa(config.Cfg.Port),
		Handler: router.Router,
	}
	logger.Logger.Printf("Server is running on port %d", config.Cfg.Port)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Logger.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	i := <-quit
	logger.Logger.Println("server receive a signal: ", i.String())

	ctx, canecl := context.WithTimeout(context.Background(), 5*time.Second)

	defer canecl()
	if err := server.Shutdown(ctx); err != nil {
		logger.Logger.Fatalf("Server failed to shutdown: %v", err)
	}
	logger.Logger.Println("Server shutdown")
}
