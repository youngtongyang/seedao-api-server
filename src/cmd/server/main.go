package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/seedao/seedao-api-server/internal/app"
	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/redisstore"
)

func defaultConfigPath() string {
	if env := os.Getenv("CONFIG_PATH"); env != "" {
		return env
	}
	candidates := []string{
		"config/config.local.yaml",
		"config/config.yaml",
		"src/config/config.local.yaml",
		"config.yaml",
	}
	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	return filepath.Join("config", "config.local.yaml")
}

func main() {
	cfg, err := config.Load(defaultConfigPath())
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	db, err := model.NewDB(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("connect db: %v", err)
	}

	redis := redisstore.New(cfg.Redis)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := redis.Ping(ctx); err != nil {
		cancel()
		log.Fatalf("connect redis: %v", err)
	}
	cancel()

	e, err := app.NewEcho(cfg, db, redis)
	if err != nil {
		log.Fatalf("init app: %v", err)
	}

	addr := cfg.Server.Addr
	go func() {
		log.Printf("seedao-api-server listening on %s", addr)
		if err := e.Start(addr); err != nil {
			log.Printf("server stopped: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown: %v", err)
	}
	_ = redis.Close()
	fmt.Println("bye")
}
