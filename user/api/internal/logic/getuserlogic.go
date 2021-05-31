package logic

import (
	"context"
	"errors"

	"zerotest/comment/rpc/commentclient"
	"zerotest/user/api/internal/svc"
	"zerotest/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req types.UserReq) (*types.UserReply, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.CommentRpc.GetComment(l.ctx, &commentclient.IdRequest{
		Id: "1",
	})
	if err != nil {
		return nil, err
	}

	if user.Name != "test" {
		return nil, errors.New("用户不存在")
	}

	return &types.UserReply{
		Id:   req.Id,
		Name: "test order",
	}, nil
}
