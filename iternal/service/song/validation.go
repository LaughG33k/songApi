package song

import "github.com/LaughG33k/songApi/iternal/model"

func validateEmptyDetailInfo(info model.DetailInfo) error {

	if info.Text == "" {
		return model.EmptyText
	}

	if info.Link == "" {
		return model.EmptyLink
	}

	if info.RealeseDate == "" {
		return model.EmptyRealese
	}

	return nil
}

func validateEmptyHead(head model.SongHead) error {
	if head.Song == "" {
		return model.EmptySong
	}

	if head.Group == "" {
		return model.EmptyGroup
	}
	return nil
}
