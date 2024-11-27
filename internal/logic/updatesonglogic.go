package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
