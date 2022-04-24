package models

type Player struct {
	Name       string `json:"name" validate:"required"`
	Gender     string `json:"gender" validate:"required"`
	Email      string `json:"email" validate:"required"`
	AvatarPath string
}


type PlayerPDF struct {
	Name string `json:"name" validate:"required"`
}

type PlayerPDFResponse struct {
	URL string `json:"url"`
}

type WorkerPdf struct {
	Name string
	Url  string
}

type PlayerAddResponse struct {
	Status string `json:"status"`
	Cause  string `json:"cause,omitempty"`
}

type GetPlayers struct {
	Players []string `json:"players"`
}

type ChangePlayer struct {
	OldName string `json:"oldName" validate:"required"`
	NewData Player `json:"newData" validate:"required"`
}
