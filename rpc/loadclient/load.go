// Code generated by goctl. DO NOT EDIT.
// Source: load.proto

package loadclient

import (
	"context"

	"github.com/vwenkk/load/rpc/types/load"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseIDResp    = load.BaseIDResp
	BaseResp      = load.BaseResp
	BaseUUIDResp  = load.BaseUUIDResp
	Empty         = load.Empty
	GroupInfo     = load.GroupInfo
	GroupListReq  = load.GroupListReq
	GroupListResp = load.GroupListResp
	IDReq         = load.IDReq
	IDsReq        = load.IDsReq
	PageInfoReq   = load.PageInfoReq
	UUIDReq       = load.UUIDReq
	UUIDsReq      = load.UUIDsReq

	Load interface {
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		// Group management
		CreateGroup(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateGroup(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetGroupList(ctx context.Context, in *GroupListReq, opts ...grpc.CallOption) (*GroupListResp, error)
		GetGroupById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*GroupInfo, error)
		DeleteGroup(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
	}

	defaultLoad struct {
		cli zrpc.Client
	}
)

func NewLoad(cli zrpc.Client) Load {
	return &defaultLoad{
		cli: cli,
	}
}

func (m *defaultLoad) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := load.NewLoadClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

// Group management
func (m *defaultLoad) CreateGroup(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := load.NewLoadClient(m.cli.Conn())
	return client.CreateGroup(ctx, in, opts...)
}

func (m *defaultLoad) UpdateGroup(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := load.NewLoadClient(m.cli.Conn())
	return client.UpdateGroup(ctx, in, opts...)
}

func (m *defaultLoad) GetGroupList(ctx context.Context, in *GroupListReq, opts ...grpc.CallOption) (*GroupListResp, error) {
	client := load.NewLoadClient(m.cli.Conn())
	return client.GetGroupList(ctx, in, opts...)
}

func (m *defaultLoad) GetGroupById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*GroupInfo, error) {
	client := load.NewLoadClient(m.cli.Conn())
	return client.GetGroupById(ctx, in, opts...)
}

func (m *defaultLoad) DeleteGroup(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := load.NewLoadClient(m.cli.Conn())
	return client.DeleteGroup(ctx, in, opts...)
}
