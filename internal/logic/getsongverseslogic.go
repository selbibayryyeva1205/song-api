package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSongVersesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSongVersesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSongVersesLogic {
	return &GetSongVersesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSongVersesLogic) GetSongVerses(req *types.SongVersesRequest) (resp *types.SongVersesResponse, err error) {
	l.WithContext(l.ctx).Infof("Fetching song verses for SongID: %d, VerseNumber: %d", req.Song_id, req.VerseNumber)
	song, err := l.svcCtx.SongModel.FindOne(l.ctx, req.Song_id, req.VerseNumber)
	if err != nil {
		l.WithContext(l.ctx).Errorf("Error fetching song verses for SongID: %d, VerseNumber: %d - %v", req.Song_id, req.VerseNumber, err)
		return nil, err
	}
	l.WithContext(l.ctx).Debugf("Found song: %+v", song)
	resp = &types.SongVersesResponse{
		Id:          song.Id,
		Group:       song.GroupName,
		Song:        song.SongName,
		Link:        song.Link.String,
		Text:        song.Text,
		ReleaseDate: song.ReleaseDate.Time.Format("02.01.2006"),
	}
	l.WithContext(l.ctx).Debugf("Song verses response: %+v", resp)

	return resp, nil
}
