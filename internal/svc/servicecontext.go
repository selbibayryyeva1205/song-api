package svc

import (
	"api/internal/config"
	"api/models/song"
	"api/models/verses"
	"database/sql"
	"fmt"
)

type ServiceContext struct {
	Config     config.Config
	VerseModel verses.VersesModel
	SongModel  song.SongsModel
	Db         sql.DB
}

func NewServiceContext(c config.Config, db sql.DB, v verses.VersesModel, s song.SongsModel) *ServiceContext {
	query := `CREATE TABLE if not exists songs (
		id SERIAL PRIMARY KEY,
		group_name TEXT NOT NULL,
		song_name TEXT NOT NULL,
		release_date time ,
		song_text TEXT NOT NULL,
		link TEXT NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println("ERROR IN CREATE", err)
	}

	query2:=`CREATE TABLE IF NOT EXISTS verses (
		id SERIAL PRIMARY KEY,
		song_id INT REFERENCES songs(id) ON DELETE CASCADE,
		verse_number INT,
		song_text TEXT NOT NULL
	
	);
	`
		_, err = db.Exec(query2)
		if err!=nil{
			fmt.Println("ERROR IN CREATE",err)
		}

	return &ServiceContext{
		Config:     c,
		VerseModel: v,
		SongModel:  s,
	}
}
