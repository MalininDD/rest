package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"rest/internal/player/delivery/http"
	"rest/internal/player/usecase"
)

func (s *Server) MapHandlers(ctx context.Context, app *fiber.App) error {
	playerUS := usecase.NewPlayerUseCase(s.cfg)
	playerHandlers := http.NewPlayerHandlers(s.cfg,playerUS)
	playerGroup := app.Group("player")
	http.MapPlayerRoutes(playerGroup, playerHandlers)

	return nil
}