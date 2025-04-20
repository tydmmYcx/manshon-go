package api

import "manshon-go/internal/api/server"

type ApiGroup struct {
	ServerApiGroup server.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
