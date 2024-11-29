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
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

func main() {
	// Load configuration from .env
	cfg := config.LoadConfig()

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", cfg.DB_DSN)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	verse := verses.NewVersesModel(sqlx.NewSqlConnFromDB(db))
	song := song.NewSongsModel(sqlx.NewSqlConnFromDB(db))
	// Initialize service context
	ctx := svc.NewServiceContext(*cfg, *db, verse, song)

	// Create the REST server
	server := rest.MustNewServer(rest.RestConf{
		Port: cfg.Port,
		Host: "127.0.0.1",
	}, rest.WithCors("*"))

	

	// RegiHster API handlers
	handler.RegisterHandlers(server, ctx)
	//server.Use(handler.CorsMiddleware(http.HandlerFu))
	//handle
	server.Use(rest.ToMiddleware(handler.CorsMiddleware))
	// Start the server
	log.Printf("Server started on port %s", cfg.Port)
	server.Start()
}
