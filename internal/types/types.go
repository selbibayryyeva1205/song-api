package types

type AddSongRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type SongActionResponse struct {
	Message string `json:"message"`
}

type SongFilterRequest struct {
	Group string `form:"group"` // Фильтрация по имени группы
	Song  string `form:"song"`  // Фильтрация по названию песни
	Page  int64  `form:"page"`  // Номер страницы
	Limit int64  `form:"limit"` // Размер страницы
}

type SongListResponse struct {
	Songs *[]Song `json:"songs"` // Список песен
	Total int64   `json:"total"` // Общее количество записей
}

type SongVersesRequest struct {
	Song_id     int64 `form:"song_id"` // ID песни
	VerseNumber int   `json:"verse_number"`
}

type SongDeleteRequest struct {
	Song_id int64 `form:"song_id"` // ID песни
}

type SongVersesResponse struct {
	Id          int64  `json:"id"`          // Уникальный идентификатор
	Group       string `json:"group"`       // Имя группы
	Song        string `json:"song"`        // Название песни
	ReleaseDate string `json:"releaseDate"` // Дата релиза
	Link        string `json:"link"`        // Ссылка на источник
	Text        string `json:"text"`
}

type Song struct {
	Id          int64  `json:"id"`          // Уникальный идентификатор
	Group       string `json:"group"`       // Имя группы
	Song        string `json:"song"`        // Название песни
	ReleaseDate string `json:"releaseDate"` // Дата релиза
	Link        string `json:"link"`        // Ссылка на источник
	Text        string `json:"text"`
}
