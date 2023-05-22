// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero-demo/greet/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: GreetHandler(serverCtx),
			},
		},
	)
}