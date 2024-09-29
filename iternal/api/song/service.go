package song

import (
	"time"

	"github.com/LaughG33k/songApi/iternal/service"
)

type Song struct {
	songService      service.SongService
	operationTimeout time.Duration
}

func NewSongApi(songService service.SongService) *Song {
	return &Song{
		songService:      songService,
		operationTimeout: time.Second * 45,
	}
}
