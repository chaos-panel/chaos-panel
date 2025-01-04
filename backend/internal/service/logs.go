// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/chaos-panel/chaos-panel/internal/model/do"
)

type (
	ILogs interface {
		Create(ctx context.Context, log *do.Logs) error
	}
)

var (
	localLogs ILogs
)

func Logs() ILogs {
	if localLogs == nil {
		panic("implement not found for interface ILogs, forgot register?")
	}
	return localLogs
}

func RegisterLogs(i ILogs) {
	localLogs = i
}
