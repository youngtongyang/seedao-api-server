package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/legacyimport"
	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/aliyunoss"
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

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: v1import <command> [flags]

Commands:
  profiles   Import 1.0 profiles/assets/roles/seeds/sbts from sns_chain_registry
  rewards    Import 1.0 new_reward applications
  oss        Backfill OSS URLs for pending seed/sbt images

Global flags:
  -config    path to YAML config (default: CONFIG_PATH or config.local.yaml)

profiles flags:
  -upload-oss   upload images during import (default: false)

oss flags:
  -limit        max rows per table per run (default: 500)

Examples:
  go run ./cmd/v1import profiles -config config/config.local.yaml
  go run ./cmd/v1import rewards -config config/config.local.yaml
  go run ./cmd/v1import oss -config config/config.local.yaml -limit 200
`)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	cmd := os.Args[1]
	fs := flag.NewFlagSet(cmd, flag.ExitOnError)
	configPath := fs.String("config", defaultConfigPath(), "config yaml path")
	uploadOSS := fs.Bool("upload-oss", false, "upload images while importing profiles")
	limit := fs.Int("limit", 500, "oss backfill batch size per table")

	args := os.Args[2:]
	if err := fs.Parse(args); err != nil {
		log.Fatal(err)
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	db, err := model.NewDB(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("connect db: %v", err)
	}

	ossClient, err := aliyunoss.NewClient(cfg.AliyunOSS)
	if err != nil {
		log.Fatalf("oss client: %v", err)
	}

	im := legacyimport.New(db, ossClient, cfg.V1Import)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	switch cmd {
	case "profiles":
		if err := im.ImportProfiles(ctx, *uploadOSS); err != nil {
			log.Fatalf("profiles: %v", err)
		}
	case "rewards":
		if err := im.ImportRewards(ctx); err != nil {
			log.Fatalf("rewards: %v", err)
		}
	case "oss":
		if err := im.BackfillOSS(ctx, *limit); err != nil {
			log.Fatalf("oss: %v", err)
		}
	case "help", "-h", "--help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", cmd)
		usage()
		os.Exit(2)
	}
}
