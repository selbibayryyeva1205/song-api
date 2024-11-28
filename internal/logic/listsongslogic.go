package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSongsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListSongsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSongsLogic {
	return &ListSongsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListSongsLogic) ListSongs(req *types.SongFilterRequest) (*types.SongListResponse, error) {
	logx.WithContext(l.ctx).Infof("Received ListSongs request: %+v", req)
	
	songs, err := l.svcCtx.SongModel.FindAll(l.ctx, &req.Group, &req.Song, int(req.Page), int(req.Limit))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Error fetching songs from database: %v", err)
		return nil, err
	}
	logx.WithContext(l.ctx).Debugf("Fetched %d songs from database", len(songs.Songs))

	var songArray []types.Song
	for _, v := range songs.Songs {
		song := &types.Song{
			Song:        v.SongName,
			Group:       v.GroupName,
			ReleaseDate: v.ReleaseDate.String(),
			Text:        v.Link.String,
		}
		logx.WithContext(l.ctx).Debugf("Processing song: %+v", song)
		songArray = append(songArray, *song)
	}

	// Build the response
	resp := &types.SongListResponse{
		Songs: &songArray,
		Total: int64(songs.Total),
	}
	logx.WithContext(l.ctx).Infof("ListSongs response prepared: Total Songs: %d", songs.Total)

	return resp, nil
}
