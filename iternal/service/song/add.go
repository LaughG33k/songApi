package song

import (
	"context"

	musicinfo "github.com/LaughG33k/songApi/iternal/client/musicInfo"
	"github.com/LaughG33k/songApi/iternal/model"
	p "github.com/LaughG33k/songApi/pkg"
)

func (s *song) Add(ctx context.Context, head model.SongHead) error {

	p.Log.Info("начинается добавления новой песни")
	p.Log.Debugf("аргументы\nНазвание: %s, Группа: %s", head.Song, head.Group)

	if err := validateEmptyHead(head); err != nil {
		p.Log.Info("добавление новой песни завершилось с ошибкой валидации")
		return err
	}

	song := model.Song{
		SongHead: head,
	}

	p.Log.Info("запрос на полуение детализированой информации")
	detailInfo, err := musicinfo.GetDetailInfo(ctx, head.Song, head.Group)

	if err != nil {
		p.Log.Info("запрос на получение детализированой закончился с ошибкой")
		return err
	}
	p.Log.Info("запрос на получение детализированой завершился успешно")

	song.DetailInfo = detailInfo
	p.Log.Debugf("детальная информация песни\n Длина Текста: %d, Сылка: %s, Релиз: %s", len(detailInfo.Text), detailInfo.Link, detailInfo.RealeseDate)

	p.Log.Info("запрос в базу на создание новой песни")
	if err := s.songRepo.Create(ctx, song); err != nil {
		p.Log.Info("запрос в базу на создание новой песни завершился с ошибкой")
		return err
	}
	p.Log.Info("запрос в базу на создание новой песни завершился успешно")

	p.Log.Info("добавление новой песни завершилось успешно")
	return nil
}
