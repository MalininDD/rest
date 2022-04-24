package server

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"os"
	"os/signal"
	"rest/config"
	"rest/pkg/httpErrors"
	"syscall"
)

const (
	ctxTimeout     = 5
	maxHeaderBytes = 1 << 20
)

type Server struct {
	fiber *fiber.App
	cfg   *config.Config
	pgDB  *sqlx.DB
}

func NewServer(
	cfg *config.Config,
) *Server {
	return &Server{
		fiber: fiber.New(fiber.Config{ErrorHandler: httpErrors.Handler, DisableStartupMessage: true}),
		cfg:   cfg,
	}
}

func (s *Server) Run(ctx context.Context) error {

	if err := s.MapHandlers(ctx, s.fiber); err != nil {
		fmt.Println("Cannot map delivery: ", err)
	}

	fmt.Println("Start server on port: %s:%s", s.cfg.Server.Host, s.cfg.Server.Port)
	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)); err != nil {
		fmt.Println("Error starting Server: ", err)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	return s.fiber.Shutdown()
}
