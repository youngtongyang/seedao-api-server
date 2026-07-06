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

// @title           SeeDAO API Server
// @version         2.7
// @description     SeeDAO 后端 API（认证、支付、课程、治理提案等）
// @host            api.seedao.xyz
// @BasePath        /
//
// @tag.name        health
// @tag.description 健康检查
// @tag.name        auth
// @tag.description 认证 /v1/auth
// @tag.name        sns
// @tag.description SNS /v1/sns
// @tag.name        wallet
// @tag.description 钱包 SIWE /v1/wallet
// @tag.name        stats
// @tag.description 统计 /v1/stats
// @tag.name        user
// @tag.description 用户 /v1/user
// @tag.name        payments
// @tag.description 支付 /v1/payments
// @tag.name        courses
// @tag.description 课程 /v1/courses
// @tag.name        governance
// @tag.description 治理配置 /v1/governance
// @tag.name        proposals
// @tag.description 提案 /v1/proposals
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description     JWT Bearer token，格式：Bearer {token}

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
