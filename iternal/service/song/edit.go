package song

import (
	"context"

	"github.com/LaughG33k/songApi/iternal/model"
	p "github.com/LaughG33k/songApi/pkg"
)

func (s *song) Edit(ctx context.Context, head model.SongHead, editFields model.Song) error {

	p.Log.Info("начало редактирования песни")
	p.Log.Debugf("аргументы для фильтрации\n Название: %s, Группа: %s", head.Song, head.Group)
	p.Log.Debugf("аргументы для редактирования\n Длина текста: %d, Ссылка: %s, Релиз: %s", len(editFields.Text), editFields.Link, editFields.RealeseDate)

	if err := validateEmptyHead(head); err != nil {
		p.Log.Info("редактирование песни завершилось с ошибкой валидаци")
		return err
	}

	p.Log.Info("запрос на редактирование песни в базе")
	if err := s.songRepo.Edit(ctx, head, editFields); err != nil {
		p.Log.Info("запрос на редактирование песни в базе завершился с ошибкой")
		return err
	}
	p.Log.Info("запрос на редактирование песни в базе завершился успешно")

	p.Log.Info("редактирование песни завершилось успешно")
	return nil
}
