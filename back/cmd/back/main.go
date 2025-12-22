package main

import (
	"back/internal/app"
	"back/internal/config"
	"back/internal/db"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Load()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.BackPostgresUser,
		cfg.BackPostgresPassword,
		cfg.BackPostgresHost,
		cfg.BackPostgresPort,
		cfg.BackPostgresDB,
	)
	conn := db.Init(dsn)
	defer db.Close(conn)
	logrus.Debug("Connected to database:", conn != nil)
	server := app.NewServer(cfg, conn)
	server.Start()
}
