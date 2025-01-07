// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IWorkerId interface {
		InitOrPanic(ctx context.Context, extra ...interface{})
		GetWorkerId(ctx context.Context) uint
	}
)

var (
	localWorkerId IWorkerId
)

func WorkerId() IWorkerId {
	if localWorkerId == nil {
		panic("implement not found for interface IWorkerId, forgot register?")
	}
	return localWorkerId
}

func RegisterWorkerId(i IWorkerId) {
	localWorkerId = i
}
