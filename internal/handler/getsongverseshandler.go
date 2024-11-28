package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"api/internal/logic"
	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSongVersesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		queryParams := r.URL.Query()
		songIDStr := queryParams.Get("song_id")
		verseNumberStr := queryParams.Get("verse_number")

		songID, err := strconv.Atoi(songIDStr)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("invalid song_id"))
			return
		}

		verseNumber, err := strconv.Atoi(verseNumberStr) // Convert verse_number to int
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("invalid verse_number"))
			return
		}

		fmt.Println("Song ID:", songID)
		fmt.Println("Verse Number:", verseNumber)
		var req *types.SongVersesRequest
		req = &types.SongVersesRequest{
			Song_id:     int64(songID),
			VerseNumber: verseNumber,
		}
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println("Error parsing request:", err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// Proceed with logic handling
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
