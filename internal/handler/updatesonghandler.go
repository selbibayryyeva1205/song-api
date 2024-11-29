package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"api/internal/logic"
	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateSongHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SongUpdate
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("err: ", err)
		}

		err = json.Unmarshal(body, &req)
		if err != nil {
			fmt.Println("err: ", err)
		}

		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) < 4 || pathParts[2] == "" {
			http.Error(w, "Song ID is required", http.StatusBadRequest)
			return
		}
		songIDStr := pathParts[3]
		songID, err := strconv.Atoi(songIDStr)
		if err != nil {
			fmt.Println("Error converting song_id to int:", err)
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("invalid song_id"))
			return
		}
		req.Id = int64(songID)
		l := logic.NewUpdateSongLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSong(&req, int64(songID))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
