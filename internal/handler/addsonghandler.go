package handler

import (
	"fmt"
	"net/http"

	"api/internal/logic"
	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
) //

func AddSongHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("TEEEST")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow specific methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// Allow specific headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		var req types.AddSongRequest
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println("1116666", err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewAddSongLogic(r.Context(), svcCtx)
		resp, err := l.AddSong(&req)
		if err != nil {
			fmt.Println("111", err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			fmt.Println("222", err)
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
