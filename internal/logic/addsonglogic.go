package logic

import (
	"context"
	"fmt"
	"log"
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
	parsedDate, err := time.Parse("", "")
	if err != nil {
		log.Fatal(err)
	}
	text := "The morning sun begins to rise\nA golden glow lights up the skies\nBut in the distance, a storm brews near\nAnd whispers of danger, I can hear\n\nThe wind it howls through the trees\nCarrying secrets, across the seas\nI try to run, but it follows me\nA force unseen, that's always free\n\nNight falls, and the world grows still\nThe wind it whispers, against my will\nI listen close, but can't make a sound\nThe whispers are lost, without a ground"
	song := &song.Songs{
		GroupName:   req.Group,
		SongName:    req.Group,
		ReleaseDate: parsedDate,
		Text:        text,
	}

	verse := strings.Split(text, "\n\n")
	res, err := l.svcCtx.SongModel.Insert(l.ctx, song)
	if err != nil {
		fmt.Println("errror", err)
	}
	//res.LastInsertId()

	fmt.Println("SONG ID ", res)
	for i, v := range verse {
		_, err := l.svcCtx.VerseModel.Insert(l.ctx, &verses.Verses{
			VerseNumber: i+1,
			SongId:      int(res),
			Text:        v,
		})
		if err != nil {
			logx.WithContext(l.ctx).Errorf("svcCtx.VerseModel.Insert((%v):", err)
		}

	}

	return
}
