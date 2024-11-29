package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"api/internal/logic"
	"api/internal/svc"

	"api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteSongHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) < 4 || pathParts[2] == "" {
			http.Error(w, "Song ID is required", http.StatusBadRequest)
			return
		}
		songIDStr := pathParts[3]
		songID, err := strconv.Atoi(songIDStr)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("invalid song_id"))
			return
		}
		var req *types.SongDeleteRequest
		req = &types.SongDeleteRequest{
			Song_id: int64(songID),
		}

		l := logic.NewDeleteSongLogic(r.Context(), svcCtx)
		resp, err := l.DeleteSong(req.Song_id)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
