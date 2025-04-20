package router

import "manshon-go/internal/router/server"

type RouterGroup struct {
	Server server.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
