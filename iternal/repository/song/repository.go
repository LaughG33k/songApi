package song

import (
	"context"
	"fmt"
	"strings"

	"github.com/LaughG33k/songApi/iternal/client/psql"
	"github.com/LaughG33k/songApi/iternal/model"
	"github.com/LaughG33k/songApi/iternal/repository"
	"github.com/jackc/pgx"
)

type song struct {
	client psql.PsqlClient
}

// Create implements repository.SongRepository.
func (s *song) Create(ctx context.Context, song model.Song) error {

	row := s.client.QueryRowEx(ctx, "select  from songs where groups = $1 and song = $2;", nil, song.Group, song.Song)

	if err := row.Scan(); err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			if _, err := s.client.ExecEx(ctx, "insert into songs(song, groups, song_text, link, realese_date) values($1, $2, $3, $4, $5);", nil, song.Song, song.Group, song.Text, song.Link, song.RealeseDate); err != nil {
				return err
			}
			return nil
		}

		return err
	}

	return model.SongExists
}

// Delete implements repository.SongRepository.
func (s *song) Delete(ctx context.Context, head model.SongHead) error {

	if _, err := s.client.ExecEx(ctx, "delete from songs where song = $1 and groups = $2;", nil, head.Song, head.Group); err != nil {
		return err
	}

	return nil
}

// Edit implements repository.SongRepository.
func (s *song) Edit(ctx context.Context, filter model.SongHead, editFields model.Song) error {

	tmpSong := filter.Song
	tmpGroup := filter.Group

	if editFields.Group != "" && editFields.Song != "" {
		tmpSong = editFields.Song
		tmpGroup = editFields.Group
	} else if editFields.Group != "" && editFields.Song == "" {
		tmpGroup = editFields.Group
	} else if editFields.Song != "" && editFields.Group == "" {
		tmpSong = editFields.Song
	}

	row := s.client.QueryRowEx(ctx, "select from songs where groups = $1 and song = $2;", nil, tmpGroup, tmpSong)

	if err := row.Scan(); err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			goto Edit
		} else {
			return err
		}
	}

	return model.SongExists

Edit:
	q := "update songs set"
	args := 1
	argsVal := make([]any, 0, 8)

	fn := func(field string, compareVal any, exceptRes any) {

		if args == 1 && compareVal != exceptRes {
			q += fmt.Sprintf(" %s = $%d", field, args)
			args++
			argsVal = append(argsVal, compareVal)
		} else if compareVal != exceptRes {
			q += fmt.Sprintf(", %s = $%d", field, args)
			args++
			argsVal = append(argsVal, compareVal)
		}

	}

	fn("song", editFields.Song, "")
	fn("groups", editFields.Group, "")
	fn("song_text", editFields.Text, "")
	fn("link", editFields.Link, "")
	fn("realese_date", editFields.RealeseDate, "")

	q += fmt.Sprintf(" where song = $%d and groups = $%d;", args, args+1)
	argsVal = append(argsVal, filter.Song)
	argsVal = append(argsVal, filter.Group)

	if _, err := s.client.ExecEx(ctx, q, nil, argsVal...); err != nil {
		return err
	}

	return nil
}

// Get implements repository.SongRepository.
func (s *song) Get(ctx context.Context, filter model.Song, limit int, offset int) ([]model.Song, error) {

	q := "select id, song, groups, song_text, realese_date, link from songs where"
	args := 1
	argsVal := make([]any, 0, 8)

	fn := func(field string, compareVal any, exceptRes any) {

		if args == 1 && compareVal != exceptRes {
			q += fmt.Sprintf(" %s = $%d", field, args)
			args++
			argsVal = append(argsVal, compareVal)
		} else if compareVal != exceptRes {
			q += fmt.Sprintf(" and %s = $%d", field, args)
			args++
			argsVal = append(argsVal, compareVal)
		}

	}

	fn("song", filter.Song, "")
	fn("groups", filter.Group, "")
	fn("song_text", filter.Text, "")
	fn("link", filter.Link, "")
	fn("realese_date", filter.RealeseDate, "")

	if args == 1 && offset > 0 {
		q += fmt.Sprintf(" id > $%d", args)
		args++
		argsVal = append(argsVal, offset)
	} else if offset > 0 {
		q += fmt.Sprintf(" and id > $%d", args)
		args++
		argsVal = append(argsVal, offset)
	}

	q += fmt.Sprintf(" order by id limit $%d;", args)
	argsVal = append(argsVal, limit)

	rows, err := s.client.QueryEx(ctx, q, nil, argsVal...)

	if err != nil {
		return nil, err
	}

	res := make([]model.Song, 0)

	for rows.Next() {

		song := model.Song{}

		if err := rows.Scan(&song.Id, &song.Song, &song.Group, &song.Text, &song.RealeseDate, &song.Link); err != nil {
			return nil, err
		}

		res = append(res, song)

	}

	return res, nil
}

// GetText implements repository.SongRepository.
func (s *song) GetText(ctx context.Context, filter model.SongHead, limit int, offset int) ([]string, error) {

	text := ""

	row := s.client.QueryRowEx(ctx, "select song_text from songs where song = $1 and groups = $2;", nil, filter.Song, filter.Group)

	if err := row.Scan(&text); err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, model.SongNotFound
		}

		return nil, err
	}

	textSplit := strings.Split(text, "\n")

	if limit > len(textSplit) {
		limit = len(textSplit)
	}

	if offset < 0 {
		offset = 0
	}

	end := 0
	if offset+limit > len(textSplit) {
		end = len(textSplit)
	} else {
		end = offset + limit
	}

	return textSplit[offset:end], nil

}

func NewRepository(client psql.PsqlClient) repository.SongRepository {
	return &song{
		client: client,
	}
}
