package logic

import (
	"context"

	"zerotest/comment/rpc/comment"
	"zerotest/comment/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *comment.IdRequest) (*comment.CommentResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.CommentResponse{}, nil
}
