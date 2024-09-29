package song

import (
	"context"

	"github.com/LaughG33k/songApi/iternal/model"
	p "github.com/LaughG33k/songApi/pkg"
)

func (s *song) GetAll(ctx context.Context, filter model.Song, limit int, offset int) ([]model.Song, error) {

	p.Log.Info("начало получения песен")
	p.Log.Debugf("аргументы для фильтрации\n Название: %s, Группа: %s, Длина текста: %d, Ссылка: %s, Релиз: %s", filter.Song, filter.Group, len(filter.Text), filter.Link, filter.RealeseDate)
	p.Log.Debugf("аргументы для пагинации\n Лимит: %d, Начало: %d", limit, offset)

	p.Log.Info("запрос на получение песен в базу")
	songs, err := s.songRepo.Get(ctx, filter, limit, offset)

	if err != nil {
		p.Log.Info("запрос на получение песен в базу зваершился с ошибкой")
		return nil, err
	}
	p.Log.Info("запрос на получение песен в базу завершился успешно")
	p.Log.Info("получение песен завершилось успешно")

	return songs, nil
}
