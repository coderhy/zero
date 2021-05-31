package logic

import (
	"context"

	"zerotest/comment/rpc/comment"
	"zerotest/comment/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentLogic) GetComment(in *comment.IdRequest) (*comment.CommentResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.CommentResponse{
		Id:     "1",
		Name:   "test rpc",
		Gender: "男",
	}, nil
}
