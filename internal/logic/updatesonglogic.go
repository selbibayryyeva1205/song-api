package logic

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"api/internal/svc"
	"api/internal/types"
	"api/models/song"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSongLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSongLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSongLogic {
	return &UpdateSongLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSongLogic) UpdateSong(req *types.Song) (resp *types.SongActionResponse, err error) {
	parsedDate, err := time.Parse("2006-01-02", req.ReleaseDate)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Error parsing release date: %v", err)
		return nil, fmt.Errorf("invalid release date format")
	}

	song := &song.SongsUpdate{
		GroupName:   req.Group,
		SongName:    req.Song,
		ReleaseDate: parsedDate,
		Text:        req.Text,
		Link: sql.NullString{
			String: req.Link,
			Valid:  req.Link != "",
		},
	}

	logx.WithContext(l.ctx).Infof("Updating song with group %s and song name %s", req.Group, req.Song)

	err = l.svcCtx.SongModel.Update(l.ctx, song)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Error updating song in database: %v", err)
		return nil, fmt.Errorf("failed to update song: %v", err)
	}

	logx.WithContext(l.ctx).Infof("Successfully updated song: %s", req.Song)
	return &types.SongActionResponse{
		Message: fmt.Sprintf("Song '%s' updated successfully", req.Song),
	}, nil
}
