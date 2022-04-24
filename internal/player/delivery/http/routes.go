package http

import (
	"github.com/gofiber/fiber/v2"
	"rest/internal/player"
)

func MapPlayerRoutes(playerRoutes fiber.Router, h player.Handlers) {
	playerRoutes.Post("/create_player", h.AddPlayer())
	playerRoutes.Get("/get_players_info", h.GetInfoPlayers())
	playerRoutes.Delete("/delete_players", h.DeletePlayers())
	playerRoutes.Put("/change_player_info", h.ChangePlayer())
	playerRoutes.Get("/get_player_info_pdf", h.GetInfoPDF())
	playerRoutes.Get("/get_player_info_pdf_finish/*", h.GetFinalPDF())
}
