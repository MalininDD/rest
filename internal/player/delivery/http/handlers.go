package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"rest/config"
	"rest/internal/models"
	"rest/internal/player"
	"rest/pkg/utils"
)

type playerHandlers struct {
	cfg      *config.Config
	playerUC player.UseCase
}

func NewPlayerHandlers(cfg *config.Config, playerUC player.UseCase) player.Handlers {
	return &playerHandlers{cfg: cfg, playerUC: playerUC}
}

func (p *playerHandlers) AddPlayer() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var params models.Player
		if err := utils.ReadRequest(c, &params); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		file, err := c.FormFile("avatar")
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(fiber.StatusBadRequest)
		}

		clientPath := uuid.New().String()
		txtPath := fmt.Sprintf("./avatars/%s.png", clientPath)
		if err := c.SaveFile(file, fmt.Sprintf("%s", txtPath)); err != nil {
			fmt.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		params.AvatarPath = clientPath
		result, err := p.playerUC.CreatePlayer(
			c.Context(),
			params)
		if err != nil {
			c.Locals("error", err.Error())
			return err
		}
		return c.JSON(result)
	}
}

func (p *playerHandlers) GetInfoPlayers() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var params models.GetPlayers
		if err := utils.ReadRequest(c, &params); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		result, err := p.playerUC.GetPlayersInfo(
			c.Context(),
			params)
		if err != nil {
			c.Locals("error", err.Error())
			return err
		}
		return c.JSON(result)
	}
}

func (p *playerHandlers) GetInfoPDF() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var params models.PlayerPDF
		if err := utils.ReadRequest(c, &params); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		result, err := p.playerUC.GetPDFInfo(
			c.Context(),
			params.Name)
		if err != nil {
			c.Locals("error", err.Error())
			return err
		}
		return c.JSON(result)
	}
}

func (p *playerHandlers) GetFinalPDF() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println(string(c.Request().URI().QueryString()))
		result := p.playerUC.GetInfoPDFFinal(
			string(c.Request().URI().QueryString()))
		//if err != nil {
		//	c.Locals("error", err.Error())
		//	return err
		//}
		return c.JSON(result)
	}
}


func (p *playerHandlers) DeletePlayers() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var params models.GetPlayers
		if err := utils.ReadRequest(c, &params); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		result, err := p.playerUC.DeletePlayers(
			c.Context(),
			params)
		if err != nil {
			c.Locals("error", err.Error())
			return err
		}
		return c.JSON(result)
	}
}

func (p *playerHandlers) ChangePlayer() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var params models.ChangePlayer
		if err := utils.ReadRequest(c, &params); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		result, err := p.playerUC.ChangePlayer(
			c.Context(),
			params)
		if err != nil {
			c.Locals("error", err.Error())
			return err
		}
		return c.JSON(result)
	}
}
