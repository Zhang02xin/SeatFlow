package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Zhang02xin/SeatFlow/internal/config"
	"github.com/Zhang02xin/SeatFlow/internal/handler"
	"github.com/Zhang02xin/SeatFlow/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		slog.Error("服务退出", "error", err)
		os.Exit(1)
	}
}

func run() error {
	cfg := config.Load()
	if cfg.MySQLDSN != "" {
		db, err := repository.OpenMySQL(cfg.MySQLDSN)
		if err != nil {
			return err
		}
		pool, err := db.DB()
		if err != nil {
			return fmt.Errorf("获取 MySQL 连接池: %w", err)
		}
		defer func() {
			if err := pool.Close(); err != nil {
				slog.Error("关闭 MySQL 连接池失败", "error", err)
			}
		}()
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/healthz", handler.Healthz)

	server := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	slog.Info("启动 HTTP 服务", "addr", cfg.HTTPAddr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("运行 HTTP 服务: %w", err)
	}
	return nil
}
