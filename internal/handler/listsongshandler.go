package handler

import (
	"net/http"
	"strconv"

	"api/internal/logic"
	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListSongsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SongFilterRequest
		// if err := httpx.Parse(r, &req); err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// 	return
		// }
		req.Group = r.URL.Query().Get("group")
		req.Song = r.URL.Query().Get("song")
		limitStr := r.URL.Query().Get("limit")
		pageSTR := r.URL.Query().Get("page")
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			http.Error(w, "Invalid or missing 'limit' parameter", http.StatusBadRequest)
			return
		}
		page, err := strconv.Atoi(pageSTR)
		if err != nil || page <= 0 {
			http.Error(w, "Invalid or missing 'page' parameter", http.StatusBadRequest)
			return
		}
		req.Limit = int64(limit)
		req.Page = int64(page)
		l := logic.NewListSongsLogic(r.Context(), svcCtx)
		resp, err := l.ListSongs(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
