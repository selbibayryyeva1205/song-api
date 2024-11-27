package svc

import (
	"api/internal/config"
	"api/models/song"
	"api/models/verses"
	"database/sql"
)

type ServiceContext struct {
	Config     config.Config
	VerseModel verses.VersesModel
	SongModel  song.SongsModel
	Db         sql.DB
}

func NewServiceContext(c config.Config, db sql.DB, v verses.VersesModel, s song.SongsModel) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		VerseModel: v,
		SongModel:  s,
	}
}
