package types

type AddSongRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type SongActionResponse struct {
	Message string `json:"message"`
}

type SongFilterRequest struct {
	Group string `form:"group"`
	Song  string `form:"song"`
	Page  int64  `form:"page" validate:"required,min=1"` // Page is now required and must be >= 1
	Limit int64  `form:"limit" validate:"required,min=1"`
}

type SongListResponse struct {
	Songs *[]Song `json:"songs"`
	Total int64   `json:"total"`
}

type SongVersesRequest struct {
	Song_id     int64 `form:"song_id"`
	VerseNumber int   `json:"verse_number"`
}

type SongDeleteRequest struct {
	Song_id int64 `form:"song_id"`
}

type SongVersesResponse struct {
	Id          int64  `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Link        string `json:"link"`
	Text        string `json:"text"`
}

type Song struct {
	Id          int64  `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Link        string `json:"link"`
	Text        string `json:"text"`
}

type SongUpdate struct {
	Id          int64  `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Link        string `json:"link"`
	Text        string `json:"text"`
}
