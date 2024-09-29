package song

import (
	"github.com/LaughG33k/songApi/iternal/repository"
	"github.com/LaughG33k/songApi/iternal/service"
)

type song struct {
	songRepo repository.SongRepository
}

func NewService(songRepo repository.SongRepository) service.SongService {
	return &song{songRepo: songRepo}
}
