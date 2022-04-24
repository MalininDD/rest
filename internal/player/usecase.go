package player

import (
	"context"
	"rest/internal/models"
)

type UseCase interface {
	CreatePlayer(ctx context.Context, params models.Player) (response models.PlayerAddResponse, err error)
	GetPlayersInfo(ctx context.Context, params models.GetPlayers) (response []models.Player, err error)
	DeletePlayers(ctx context.Context, params models.GetPlayers) (response models.PlayerAddResponse, err error)
	ChangePlayer(ctx context.Context, params models.ChangePlayer) (response models.PlayerAddResponse, err error)
	GetPDFInfo(ctx context.Context, name string) (response models.PlayerPDFResponse, err error)
	GetInfoPDFFinal(uuid string) (response models.PlayerAddResponse)
}

