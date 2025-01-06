// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"time"
)

type (
	IDlock interface {
		Lock(ctx context.Context, key string, ttl time.Duration, by string, host string) error
		Unlock(ctx context.Context, key string, by string) error
	}
)

var (
	localDlock IDlock
)

func Dlock() IDlock {
	if localDlock == nil {
		panic("implement not found for interface IDlock, forgot register?")
	}
	return localDlock
}

func RegisterDlock(i IDlock) {
	localDlock = i
}
