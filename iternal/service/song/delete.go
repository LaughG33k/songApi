package song

import (
	"context"

	"github.com/LaughG33k/songApi/iternal/model"
	p "github.com/LaughG33k/songApi/pkg"
)

func (s *song) Delete(ctx context.Context, head model.SongHead) error {

	p.Log.Info("начало удаления песни")
	p.Log.Debugf("аргументы для удаления песни\n Название: %s, Группа: %s", head.Song, head.Group)
	if err := validateEmptyHead(head); err != nil {
		p.Log.Info("удаление песни завершилось с ошибкой валидации")
		return err
	}

	p.Log.Info("запрос на удаление песни в базе")
	if err := s.songRepo.Delete(ctx, head); err != nil {
		p.Log.Info("запрос на удаление песни в базе завершился с ошибкой")
		return err
	}
	p.Log.Info("запрос на удаление песни в базе завершился успешно")

	p.Log.Info("удаление песни завершилось успешно")
	return nil
}
