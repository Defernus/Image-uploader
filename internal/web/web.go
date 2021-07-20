package web

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"imageapi.lavrentev.dev/rest/internal/database"
)

const (
	internalError = "Internal server error"
)

type Server struct {
	cfg        Config
	db         database.Database
	log        logrus.FieldLogger
	httpServer *http.Server

	shutdown chan struct{}
}

type Config struct {
	Addr            string
	ShutdownTimeout time.Duration
}

func NewServer(
	cfg Config,
	db database.Database,
	log logrus.FieldLogger,
) (*Server, error) {
	s := &Server{
		cfg:        cfg,
		db:         db,
		log:        log,
		httpServer: nil,
		shutdown:   make(chan struct{}),
	}

	router := mux.NewRouter()

	// Middlewares
	router.Use(s.LogRequest)

	// Methods
	router.Methods("GET").Path("/").HandlerFunc(s.handleGetImages)

	s.httpServer = &http.Server{
		Addr:              cfg.Addr,
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	return s, nil
}

func (s *Server) Start() error {
	s.log.Infof("start server at http://0.0.0.0%v", s.cfg.Addr)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Shutdown() error {
	close(s.shutdown)
	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}
