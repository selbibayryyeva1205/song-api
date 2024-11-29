package logic

import (
	"context"
	"fmt"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSongLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSongLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSongLogic {
	return &DeleteSongLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSongLogic) DeleteSong(song_id int64) (resp *types.SongActionResponse, err error) {
	l.WithContext(l.ctx).Infof("Received DeleteSong request")
	err = l.svcCtx.SongModel.Delete(l.ctx, song_id)
	if err != nil {
		l.WithContext(l.ctx).Errorf("Error deleting song from database: %v", err)
		return nil, err
	}
	return &types.SongActionResponse{
		Message: fmt.Sprintf("Song with id %d  updated successfully", song_id),
	}, nil
}
