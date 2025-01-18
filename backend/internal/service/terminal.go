// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/chaos-plus/chaos-plus/api/terminals/v1"
)

type (
	ITerminal interface {
		Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
		GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error)
		GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	}
)

var (
	localTerminal ITerminal
)

func Terminal() ITerminal {
	if localTerminal == nil {
		panic("implement not found for interface ITerminal, forgot register?")
	}
	return localTerminal
}

func RegisterTerminal(i ITerminal) {
	localTerminal = i
}
