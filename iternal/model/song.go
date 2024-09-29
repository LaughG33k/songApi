package model

type SongHead struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type DetailInfo struct {
	Text        string `json:"text"`
	RealeseDate string `json:"releaseDate"`
	Link        string `json:"link"`
}

type Song struct {
	Id int `json:"id"`
	SongHead
	DetailInfo
}
