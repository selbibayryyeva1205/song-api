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

func (l *UpdateSongLogic) UpdateSong(req *types.SongUpdate, id int64) (resp *types.SongActionResponse, err error) {
	var releaseDate sql.NullTime
	if req.ReleaseDate == "" {
		releaseDate = sql.NullTime{Valid: false} // This marks the date as NULL in DB
		l.WithContext(l.ctx).Info("Release date is empty, setting as NULL in DB")
	} else {
		parsedDate, err := time.Parse("02.01.2006", req.ReleaseDate)
		if err != nil {
			l.WithContext(l.ctx).Errorf("Error parsing release date: %v", err)
			return nil, fmt.Errorf("invalid release date format")
		}
		releaseDate = sql.NullTime{Time: parsedDate, Valid: true}
		l.WithContext(l.ctx).Infof("Parsed release date: %v", parsedDate)
	}

	song := &song.SongsUpdate{
		Id: id,

		GroupName:   req.Group,
		SongName:    req.Song,
		ReleaseDate: releaseDate,
		Text:        req.Text,
		Link: sql.NullString{
			String: req.Link,
			Valid:  req.Link != "",
		},
	}

	l.WithContext(l.ctx).Infof("Updating song with group %s and song name %s", req.Group, req.Song)

	err = l.svcCtx.SongModel.Update(l.ctx, song)
	if err != nil {
		l.WithContext(l.ctx).Errorf("Error updating song in database: %v", err)
		return nil, fmt.Errorf("failed to update song: %v", err)
	}

	l.WithContext(l.ctx).Infof("Successfully updated song: %s", req.Song)
	return &types.SongActionResponse{
		Message: fmt.Sprintf("Song '%s' updated successfully", req.Song),
	}, nil
}
