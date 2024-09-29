package song

import (
	"context"

	"github.com/LaughG33k/songApi/iternal/model"
	p "github.com/LaughG33k/songApi/pkg"
)

func (s *song) GetText(ctx context.Context, filter model.SongHead, limit int, offset int) ([]string, error) {

	p.Log.Info("начало получения текста")
	p.Log.Debugf("аргументы для фильтрации\n Название: %s, Группа: %s", filter.Song, filter.Group)
	p.Log.Debugf("аргументы для пагинации\n Лимит: %d, Начало: %d", limit, offset)

	if err := validateEmptyHead(filter); err != nil {
		p.Log.Debug("получения текста завершилось с ошибкой валидации")
		return nil, err
	}

	p.Log.Info("запрос на получения текста")
	text, err := s.songRepo.GetText(ctx, filter, limit, offset)

	if err != nil {
		p.Log.Info("запрос на получения текста завершился с ошибкой")
		return nil, err
	}
	p.Log.Info("запрос на получения текста завершился успешно")
	p.Log.Info("получения текста завершилось успешно")

	return text, nil

}
