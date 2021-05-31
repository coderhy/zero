package svc

import (
	"zerotest/comment/rpc/commentclient"
	"zerotest/user/api/internal/config"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	CommentRpc commentclient.Comment
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CommentRpc: commentclient.NewComment(zrpc.MustNewClient(c.CommentRpc)),
	}
}
