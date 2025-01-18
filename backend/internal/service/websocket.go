// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IWebsocket interface {
		Upgrade(r *ghttp.Request) error
	}
)

var (
	localWebsocket IWebsocket
)

func Websocket() IWebsocket {
	if localWebsocket == nil {
		panic("implement not found for interface IWebsocket, forgot register?")
	}
	return localWebsocket
}

func RegisterWebsocket(i IWebsocket) {
	localWebsocket = i
}
