package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"api/internal/svc"
	"api/internal/types"
	"api/models/song"
	"api/models/verses"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSongLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSongLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSongLogic {
	return &AddSongLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSongLogic) AddSong(req *types.AddSongRequest) (resp *types.SongActionResponse, err error) {
	var text string
	var song2 *song.Songs
	var parsedDate time.Time

	parsedDate, err = time.Parse("2006-01-02", "2024-01-01")
	if err != nil {
		return nil, fmt.Errorf("invalid release date")
	}

	apiData, err := getSongInfo(req.Group, req.Song, l.svcCtx.Config.OpenAPI)
	if err != nil || apiData == nil {
		text = "The morning sun begins to rise\nA golden glow lights up the skies\nBut in the distance, a storm brews near\nAnd whispers of danger, I can hear\n\nThe wind it howls through the trees\nCarrying secrets, across the seas\nI try to run, but it follows me\nA force unseen, that's always free\n\nNight falls, and the world grows still\nThe wind it whispers, against my will\nI listen close, but can't make a sound\nThe whispers are lost, without a ground"
	} else {
		text = apiData.Text
		parsedDate := time.Parse("2006-01-02", apiData.ReleaseDate)
		if err != nil {
			return nil, fmt.Errorf("invalid release date format from API")
		}
	}

	song := &song.Songs{
		GroupName:   req.Group,
		SongName:    req.Song,
		ReleaseDate: parsedDate,
		Text:        text,
		Link: sql.NullString{
			String: apiData.Link,
			Valid:  apiData.Link != "",
		},
	}
	song2 = song

	res, err := l.svcCtx.SongModel.Insert(l.ctx, song2)
	if err != nil {
		return nil, fmt.Errorf("failed to insert song")
	}

	versesText := strings.Split(text, "\n\n")

	for i, v := range versesText {
		verse := &verses.Verses{
			VerseNumber: i + 1,
			SongId:      int(res),
			Text:        v,
		}
		_, err := l.svcCtx.VerseModel.Insert(l.ctx, verse)
		if err != nil {
			continue
		}
	}

	resp = &types.SongActionResponse{
		Message: fmt.Sprintf("Song added successfully with ID: %d", res),
	}
	return resp, nil
}

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func getSongInfo(group, song string, baseURL string) (*SongDetail, error) {
	params := url.Values{}
	params.Add("group", group)
	params.Add("song", song)
	url := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var songDetail SongDetail
	err = json.Unmarshal(body, &songDetail)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return &songDetail, nil
}
