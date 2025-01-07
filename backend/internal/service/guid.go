// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IGuid interface {
		InitOrPanic(ctx context.Context)
		NextId() uint64
	}
)

var (
	localGuid IGuid
)

func Guid() IGuid {
	if localGuid == nil {
		panic("implement not found for interface IGuid, forgot register?")
	}
	return localGuid
}

func RegisterGuid(i IGuid) {
	localGuid = i
}
