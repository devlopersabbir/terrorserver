package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/devlopersabbir/terrorserver/internal/logger"
	"github.com/devlopersabbir/terrorserver/internal/server"
	"github.com/devlopersabbir/terrorserver/internal/watcher"
)

const (
	defaultConfigPath = "/etc/terror/Runtime"
	defaultAddr       = ":80"
)

var version = "prod"

func main() {
	if len(os.Args) < 2 {
		runServe()
		return
	}

	switch os.Args[1] {
	case "start", "serve", "s":
		runServe()
	case "status", "st":
		runStatus()
	case "validate", "v":
		runValidate()
	case "version", "--version", "-v":
		fmt.Printf("terror version %s\n", version)
	case "help", "-h", "--help":
		printHelp()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		printHelp()
		os.Exit(1)
	}
}

func configPath() string {
	if p := os.Getenv("TERROR_CONFIG"); p != "" {
		return p
	}
	return defaultConfigPath
}

func listenAddr() string {
	if a := os.Getenv("TERROR_ADDR"); a != "" {
		return a
	}
	return defaultAddr
}

func runServe() {
	cfgPath := configPath()
	addr := listenAddr()

	logger.Info("terrorserver %s starting", version)
	logger.Info("config: %s", cfgPath)

	srv := server.New(cfgPath)

	// Initial config load — fatal if it fails on startup
	if err := srv.LoadConfig(); err != nil {
		logger.Fatal("failed to load config: %v", err)
	}

	// Start HTTP server
	if err := srv.Start(addr); err != nil {
		logger.Fatal("failed to start server: %v", err)
	}

	// Watch config file for live reload
	done := make(chan struct{})
	if err := watcher.Watch(cfgPath, done, func() {
		if err := srv.LoadConfig(); err != nil {
			logger.Error("config reload failed, keeping old config: %v", err)
		} else {
			logger.Info("config reloaded successfully")
		}
	}); err != nil {
		logger.Warn("file watcher unavailable: %v — live reload disabled", err)
	}

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down gracefully...")
	close(done)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("shutdown error: %v", err)
	}
	logger.Info("stopped")
}

func runStatus() {
	printStatus(configPath(), listenAddr())
}

func runValidate() {
	cfgPath := configPath()
	fmt.Printf("validating %s ...\n", cfgPath)

	// Import config inline to validate
	// We re-use the server's LoadConfig approach
	srv := server.New(cfgPath)
	if err := srv.LoadConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "✘ invalid config: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✔ config is valid")
}

func printHelp() {
	fmt.Printf(`terrorserver %s — minimal domain router & reverse proxy

Usage:
  terror [command]

Commands:
  start, serve   > [s] Start the server (default)
  validate       > [v] Validate the config file without starting
  status         > [st, s] Show runtime status
  version        > [-v, --version] Print version
  help           > [-h, --help] Show this help

Environment:
  TERROR_CONFIG   Path to config file (default: /etc/terror/Runtime)
  TERROR_ADDR     Listen address (default: :80)
`, version)
}
