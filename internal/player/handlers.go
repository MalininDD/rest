package player

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	AddPlayer() fiber.Handler
	GetInfoPlayers() fiber.Handler
	DeletePlayers() fiber.Handler
	ChangePlayer() fiber.Handler
	GetFinalPDF() fiber.Handler
	GetInfoPDF() fiber.Handler
}
