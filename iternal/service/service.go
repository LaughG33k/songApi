package service

import (
	"context"

	"github.com/LaughG33k/songApi/iternal/model"
)

type SongService interface {
	GetAll(ctx context.Context, filter model.Song, limit, offset int) ([]model.Song, error)
	GetText(ctx context.Context, filter model.SongHead, limit, offset int) ([]string, error)
	Delete(context.Context, model.SongHead) error
	Edit(context.Context, model.SongHead, model.Song) error
	Add(context.Context, model.SongHead) error
}
