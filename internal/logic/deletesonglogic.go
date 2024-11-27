package logic

import (
	"context"

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

func (l *DeleteSongLogic) DeleteSong() (resp *types.SongActionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
