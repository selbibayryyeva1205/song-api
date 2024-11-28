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

func (l *ListSongsLogic) ListSongs(req *types.SongFilterRequest) (resp *types.SongListResponse, err error) {
	// todo: add your logic here and delete this line
	//songs,_:=l.svcCtx.SongModel.FindAll(l.ctx,&req.Group,&req.Song,int(req.Limit),int(req.Page))

	return
}
