package svc

import (
	"api/internal/config"
	"api/models/song"
	"api/models/verses"
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config     config.Config
	VerseModel verses.VersesModel
	SongModel  song.SongsModel
	Db         *sql.DB
}

func NewServiceContext(c config.Config, db *sql.DB, v verses.VersesModel, s song.SongsModel) *ServiceContext {
	logx.Infof("Initializing ServiceContext")

	songsTableQuery := `
	CREATE TABLE IF NOT EXISTS songs (
		id SERIAL PRIMARY KEY,
		group_name TEXT,
		song_name TEXT,
		release_date DATE,
		song_text TEXT,
		link TEXT,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	`
	logx.Debugf("Executing query to create 'songs' table: %s", songsTableQuery)
	if _, err := db.Exec(songsTableQuery); err != nil {
		logx.Errorf("Failed to create 'songs' table: %v", err)
		//panic(err)
	}
	logx.Infof("'songs' table created or already exists")

	versesTableQuery := `
	CREATE TABLE IF NOT EXISTS verses (
		id SERIAL PRIMARY KEY,
		song_id INT REFERENCES songs(id) ON DELETE CASCADE,
		verse_number INT,
		song_text TEXT NOT NULL
	);
	`
	logx.Debugf("Executing query to create 'verses' table: %s", versesTableQuery)
	if _, err := db.Exec(versesTableQuery); err != nil {
		logx.Errorf("Failed to create 'verses' table: %v", err)
		//panic(err)
	}
	logx.Infof("'verses' table created or already exists")

	logx.Infof("ServiceContext initialized successfully")

	return &ServiceContext{
		Config:     c,
		VerseModel: v,
		SongModel:  s,
		Db:         db,
	}
}
