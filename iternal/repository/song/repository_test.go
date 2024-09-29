package song

import (
	"context"
	"testing"
	"time"

	"github.com/LaughG33k/songApi/iternal/client/psql"
	"github.com/LaughG33k/songApi/iternal/config"
	"github.com/LaughG33k/songApi/iternal/model"
	"github.com/LaughG33k/songApi/iternal/repository"
)

func initRepository() (repository.SongRepository, error) {

	client, err := psql.NewClient(context.TODO(), config.DBCfg{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "postgres",
		Password: "123456",
		DB:       "songdb",
	})

	if err != nil {
		return nil, err
	}

	return NewRepository(client), nil
}

func TestCreate(t *testing.T) {

	repo, err := initRepository()

	if err != nil {
		t.Error(err)
	}

	type testCase struct {
		Name        string
		Value       model.Song
		ExpectedRes any
	}

	cases := []testCase{
		testCase{
			Name: "correct create",
			Value: model.Song{
				SongHead: model.SongHead{
					Song:  "sky",
					Group: "playboi carti",
				},
				DetailInfo: model.DetailInfo{
					Text:        "What? What? What?\n What?\nI told my boy,\n Go roll\n like ten \nblunts for me \n(what? Roll ten,\n what?)\nI'm tryna get high\n 'til I can't feel nothin' (whoa, what?\n What? What?)",
					Link:        "https://genius.com/Genius-russian-translations-playboi-carti-sky-lyrics",
					RealeseDate: "25.12.2020",
				},
			},
			ExpectedRes: nil,
		},
		testCase{
			Name: "empty create",
			Value: model.Song{
				SongHead: model.SongHead{
					Song:  "skif",
					Group: "playboi carti",
				},
				DetailInfo: model.DetailInfo{
					Text:        "What? What? What? What?\nI told my boy, Go roll like ten blunts for me (what? Roll ten, what?)\nI'm tryna get high 'til I can't feel nothin' (whoa, what? What? What?)",
					Link:        "https://genius.com/Genius-russian-translations-playboi-carti-sky-lyrics",
					RealeseDate: "25.12.2020",
				},
			},
			ExpectedRes: nil,
		},
	}

	for _, v := range cases {
		val := v

		t.Run(val.Name, func(t *testing.T) {
			var tm context.Context
			var canc context.CancelFunc

			if deadline, ok := t.Deadline(); ok {
				tm, canc = context.WithDeadline(context.TODO(), deadline)
			} else {
				tm, canc = context.WithTimeout(context.TODO(), 35*time.Second)
			}

			defer canc()

			if err := repo.Create(tm, val.Value); err != nil {
				t.Error(err)
			}
		})
	}

}

func TestDelete(t *testing.T) {
	repo, err := initRepository()

	if err != nil {
		t.Error(err)
	}
	type testCase struct {
		Name        string
		Value       model.SongHead
		ExpectedRes any
	}

	cases := []testCase{
		testCase{
			Name: "default",
			Value: model.SongHead{
				Group: "playboi carti",
			},
		},
	}

	for _, v := range cases {
		val := v

		t.Run(val.Name, func(t *testing.T) {
			var tm context.Context
			var canc context.CancelFunc

			if deadline, ok := t.Deadline(); ok {
				tm, canc = context.WithDeadline(context.TODO(), deadline)
			} else {
				tm, canc = context.WithTimeout(context.TODO(), 35*time.Second)
			}

			defer canc()

			if err := repo.Delete(tm, val.Value); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestGetText(t *testing.T) {

	repo, err := initRepository()

	if err != nil {
		t.Error(err)
	}
	type testCase struct {
		Name        string
		Value       model.SongHead
		Limit       int
		Offset      int
		ExpectedRes any
	}

	cases := []testCase{
		testCase{
			Name: "getText",
			Value: model.SongHead{
				Group: "playboi carti",
				Song:  "sky",
			},
			Limit:       2,
			Offset:      10,
			ExpectedRes: nil,
		},
	}

	for _, v := range cases {
		val := v

		t.Run(val.Name, func(t *testing.T) {
			var tm context.Context
			var canc context.CancelFunc

			if deadline, ok := t.Deadline(); ok {
				tm, canc = context.WithDeadline(context.TODO(), deadline)
			} else {
				tm, canc = context.WithTimeout(context.TODO(), 35*time.Second)
			}

			defer canc()

			couplets, err := repo.GetText(tm, v.Value, v.Limit, val.Offset)

			if err != nil {
				t.Error(err)
			}

			for _, v := range couplets {
				t.Log(v)
			}

		})
	}

}

func TestEdit(t *testing.T) {

	repo, err := initRepository()

	if err != nil {
		t.Error(err)
	}

	type testCase struct {
		Name        string
		EditFields  model.Song
		Filter      model.SongHead
		ExpectedRes any
	}

	cases := []testCase{
		testCase{
			Name: "default",
			EditFields: model.Song{
				SongHead: model.SongHead{
					Group: "playboi carti",
					Song:  "skyf",
				},
				DetailInfo: model.DetailInfo{
					Text: "nety texta",
				},
			},

			Filter: model.SongHead{
				Song:  "sky",
				Group: "sky",
			},
		},
	}

	for _, v := range cases {
		val := v

		t.Run(val.Name, func(t *testing.T) {
			var tm context.Context
			var canc context.CancelFunc

			if deadline, ok := t.Deadline(); ok {
				tm, canc = context.WithDeadline(context.TODO(), deadline)
			} else {
				tm, canc = context.WithTimeout(context.TODO(), 35*time.Second)
			}

			defer canc()

			if err := repo.Edit(tm, val.Filter, val.EditFields); err != nil {
				t.Error(err)
			}

		})
	}

}

func TestGet(t *testing.T) {
	repo, err := initRepository()

	if err != nil {
		t.Error(err)
	}

	type testCase struct {
		Name        string
		Value       model.Song
		Limit       int
		Offset      int
		ExpectedRes any
	}

	cases := []testCase{
		testCase{
			Name: "get",
			Value: model.Song{

				SongHead: model.SongHead{
					Song: "sky",
				},
				DetailInfo: model.DetailInfo{
					RealeseDate: "25.12.2020",
				},
			},
			Limit:       3,
			Offset:      1,
			ExpectedRes: nil,
		},
	}

	for _, v := range cases {
		val := v
		t.Run(val.Name, func(t *testing.T) {

			var tm context.Context
			var canc context.CancelFunc

			if deadline, ok := t.Deadline(); ok {
				tm, canc = context.WithDeadline(context.TODO(), deadline)
			} else {
				tm, canc = context.WithTimeout(context.TODO(), 35*time.Second)
			}

			defer canc()

			songs, err := repo.Get(tm, val.Value, val.Limit, val.Offset)

			if err != nil {
				t.Error(err)
			}

			t.Log(songs)

		})
	}
}
