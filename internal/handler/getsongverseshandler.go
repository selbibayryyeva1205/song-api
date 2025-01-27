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

func GetSongVersesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) < 5 || pathParts[2] == "" {
			http.Error(w, "Song ID is required", http.StatusBadRequest)
			return
		}
		songIDStr := pathParts[3]
		verseNumberStr := pathParts[4]
		songID, err := strconv.Atoi(songIDStr)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("invalid song_id"))
			return
		}

		verseNumber, err := strconv.Atoi(verseNumberStr) 
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("invalid verse_number"))
			return
		}
		var req *types.SongVersesRequest
		req = &types.SongVersesRequest{
			Song_id:     int64(songID),
			VerseNumber: verseNumber,
		}
		l := logic.NewGetSongVersesLogic(r.Context(), svcCtx)
		resp, err := l.GetSongVerses(&types.SongVersesRequest{
			Song_id:     req.Song_id,
			VerseNumber: verseNumber,
		})
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
