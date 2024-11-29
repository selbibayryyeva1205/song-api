// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/songs",
				Handler: ListSongsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/songs/verse/:id/:verse_id",
				Handler: GetSongVersesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/songs/create",
				Handler: AddSongHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/songs/update/:id",
				Handler: UpdateSongHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/songs/delete/:id",
				Handler: DeleteSongHandler(serverCtx),
			},
		},
	)
}
