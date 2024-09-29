package app

import (
	"context"
	"log"

	apiSong "github.com/LaughG33k/songApi/iternal/api/song"
	"github.com/LaughG33k/songApi/iternal/client/psql"
	"github.com/LaughG33k/songApi/iternal/config"
	"github.com/LaughG33k/songApi/iternal/repository"
	songRepo "github.com/LaughG33k/songApi/iternal/repository/song"
	"github.com/LaughG33k/songApi/iternal/service"
	songSrvc "github.com/LaughG33k/songApi/iternal/service/song"
)

type serviceProvider struct {
	songRepository repository.SongRepository
	songService    service.SongService
	songApi        *apiSong.Song
	config         config.AppCfg
}

func newServiceProvider() *serviceProvider {
	p := &serviceProvider{}

	p.ConfigLoad()

	return p
}

func (p *serviceProvider) ConfigLoad() config.AppCfg {

	config, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	p.config = config

	return p.config
}

func (p *serviceProvider) SongRepository(ctx context.Context) repository.SongRepository {

	if p.songRepository == nil {

		client, err := psql.NewClient(ctx, p.config.DB)

		if err != nil {
			log.Fatal(err)
		}

		repo := songRepo.NewRepository(client)

		p.songRepository = repo
	}

	return p.songRepository

}

func (p *serviceProvider) SongService(ctx context.Context) service.SongService {

	if p.songService == nil {

		p.songService = songSrvc.NewService(p.SongRepository(ctx))

	}

	return p.songService

}

func (p *serviceProvider) SongApi(ctx context.Context) *apiSong.Song {

	if p.songApi == nil {
		p.songApi = apiSong.NewSongApi(p.SongService(ctx))
	}

	return p.songApi

}
