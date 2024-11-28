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

func DeleteSongHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get query parameters
		queryParams := r.URL.Query()

		// Example: Extract and convert specific query parameters, e.g., "song_id" and "verse_number"
		songIDStr := queryParams.Get("song_id") // Get the query parameter as a string
		verseNumberStr := queryParams.Get("verse_number")

		// Convert the string parameters to integers
		songID, err := strconv.Atoi(songIDStr) // Convert songID to int
		if err != nil {
			// Handle the error if conversion fails
			fmt.Println("Error converting song_id to int:", err)
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("invalid song_id"))
			return
		}

		verseNumber, err := strconv.Atoi(verseNumberStr) // Convert verse_number to int
		if err != nil {
			// Handle the error if conversion fails
			fmt.Println("Error converting verse_number to int:", err)
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("invalid verse_number"))
			return
		}

		// Log or handle the integer values
		fmt.Println("Song ID:", songID)
		fmt.Println("Verse Number:", verseNumber)
		var req *types.SongDeleteRequest
		req = &types.SongDeleteRequest{
			Song_id: int64(songID),
			//VerseNumber: int64(verseNumber),
		}
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println("Error parsing request:", err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// Proceed with logic handling
		l := logic.NewDeleteSongLogic(r.Context(), svcCtx)
		resp, err := l.DeleteSong(req.Song_id)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
