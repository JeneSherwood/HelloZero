package handler

import (
	"net/http"

	"go-zero-demo/book/service/search/cmd/api/internal/logic"
	"go-zero-demo/book/service/search/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func pingHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewPingLogic(r.Context(), ctx)
		err := l.Ping()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
