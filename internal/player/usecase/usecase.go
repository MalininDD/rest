package usecase

import (
	"context"
	"github.com/google/uuid"
	"rest/config"
	"rest/internal/models"
	"rest/internal/player"
)

type playerUC struct {
	cfg       *config.Config
	Players   []models.Player
	PDFWorker chan models.WorkerPdf
	PDFInfo   map[string]models.PlayerAddResponse
}

func NewPlayerUseCase(cfg *config.Config) player.UseCase {
	c := make(chan models.WorkerPdf)
	return &playerUC{
		cfg:       cfg,
		PDFWorker: c,
	}
}

func (p *playerUC) CreatePlayer(ctx context.Context, params models.Player) (response models.PlayerAddResponse, err error) {
	for _, r := range p.Players {
		if r.Name == params.Name {
			return models.PlayerAddResponse{Status: "Cannot", Cause: "Such player already exists"}, nil
		}
	}

	p.Players = append(p.Players, params)
	return models.PlayerAddResponse{Status: "Successfully added"}, nil
}

func (p *playerUC) GetPlayersInfo(ctx context.Context, params models.GetPlayers) (response []models.Player, err error) {
	for _, r := range params.Players {
		for _, l := range p.Players {
			if r == l.Name {
				response = append(response, l)
				break
			}
		}
	}
	return
}

func (p *playerUC) DeletePlayers(ctx context.Context, params models.GetPlayers) (response models.PlayerAddResponse, err error) {
	for _, r := range params.Players {
		for i, l := range p.Players {
			if r == l.Name {
				p.Players = append(p.Players[:i], p.Players[i+1:]...)
				break
			}
		}
	}

	return models.PlayerAddResponse{Status: "Successfully deleted"}, nil
}

func (p *playerUC) ChangePlayer(ctx context.Context, params models.ChangePlayer) (response models.PlayerAddResponse, err error) {
	for i, r := range p.Players {
		if r.Name == params.OldName {
			p.Players[i] = params.NewData
		}
	}

	return models.PlayerAddResponse{Status: "Successfully changed"}, nil
}

func (p *playerUC) GetPDFInfo(ctx context.Context, name string) (response models.PlayerPDFResponse, err error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return response, err
	}
	data := models.WorkerPdf{Name: name, Url: p.cfg.Server.Host + ":" + p.cfg.Server.Port + "/player" + "/get_pdf_info" + "/" + uuid.String() }
	p.PDFWorker <- data

	p.PDFInfo[uuid.String()] = models.PlayerAddResponse{Status: "Accepted"}
	return models.PlayerPDFResponse{URL: p.cfg.Server.Host + ":" + p.cfg.Server.Port + "/player" + "/get_pdf_info" + "/" + uuid.String()}, nil
}

func (p *playerUC) GetInfoPDFFinal(uuid string) (response models.PlayerAddResponse){
	if val, ok := p.PDFInfo[uuid]; ok {
		return val
	}
	return models.PlayerAddResponse{Status: "Cannot", Cause: "Not found"}
}


func (p *playerUC) WorkerPDF() {
	for {
		res, ok := <-p.PDFWorker
		if !ok {
			break
		}
		for _, r := range p.Players {
			if r.Name == res.Name {

			}
		}
	}
}
