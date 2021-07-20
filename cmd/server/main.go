package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"imageapi.lavrentev.dev/rest/internal/database"
	"imageapi.lavrentev.dev/rest/internal/web"
)

type Config struct {
	Web web.Config
}

func main() {
	cfg := Config{
		Web: web.Config{
			Addr:            ":8080",
			ShutdownTimeout: 2 * time.Second,
		},
	}

	log := logrus.New()
	log.Level = logrus.DebugLevel

	log.Info("init database")
	db, err := database.NewDatabase(log)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.StartMigrations(); err != nil {
		log.Fatal(err)
	}

	log.Info("init server")
	server, err := web.NewServer(cfg.Web, db, log)
	if err != nil {
		log.Fatal(err)
	}

	serverError := make(chan error, 1)
	go func() {
		log.Info("start the server")
		serverError <- server.Start()
	}()

	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-term:
		log.Warn("got an interrupt signal")
	case err := <-serverError:
		log.WithError(err).Warn("server is down")
	}

	log.Info("shutdown")
	if err := server.Shutdown(); err != nil {
		log.WithError(err).Error("Successfully failed to shutdown server gracefully")
	}
}
