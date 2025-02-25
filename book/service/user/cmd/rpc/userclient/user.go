// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

//go:generate mockgen -destination ./user_mock.go -package userclient -source $GOFILE

package userclient

import (
	"context"

	"go-zero-demo/book/service/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type (
	IdReq         = user.IdReq
	UserInfoReply = user.UserInfoReply

	User interface {
		GetUser(ctx context.Context, in *IdReq) (*UserInfoReply, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUser(ctx context.Context, in *IdReq) (*UserInfoReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in)
}
