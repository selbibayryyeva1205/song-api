package handler

import (
	"net/http"

	"api/internal/logic"
	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
) //

func AddSongHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddSongRequest
		if err := httpx.Parse(r, &req); err != nil {

			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewAddSongLogic(r.Context(), svcCtx)
		resp, err := l.AddSong(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
