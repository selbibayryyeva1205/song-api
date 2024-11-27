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
	// todo: add your logic here and delete this line

	return
}
