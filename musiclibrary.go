package main

import (
	"api/internal/config"
	"api/internal/handler"
	"api/internal/svc"
	"api/models/song"
	"api/models/verses"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

func main() {
	logx.Infof("Loading configuration")
	cfg := config.LoadConfig()

	logx.Infof("Connecting to PostgreSQL database")
	db, err := sql.Open("postgres", cfg.DB_DSN)
	if err != nil {
		logx.Errorf("Failed to connect to database: %v", err)
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		logx.Infof("Closing database connection")
		if err := db.Close(); err != nil {
			logx.Errorf("Error closing database connection: %v", err)
		}
	}()

	logx.Infof("Initializing models")
	verse := verses.NewVersesModel(sqlx.NewSqlConnFromDB(db))
	song := song.NewSongsModel(sqlx.NewSqlConnFromDB(db))

	logx.Infof("Initializing service context")
	ctx := svc.NewServiceContext(*cfg, db, verse, song)

	logx.Infof("Creating REST server on %s:%d", cfg.Host, cfg.Port)
	server := rest.MustNewServer(rest.RestConf{
		Port: cfg.Port,
		Host: cfg.Host,
	}, rest.WithCors("*"))

	logx.Infof("Registering API handlers")
	handler.RegisterHandlers(server, ctx)

	logx.Infof("Starting server on port %d", cfg.Port)
	server.Start()
	logx.Infof("Server has stopped")
}
