package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func reverseProxyAPI(ctx *gin.Context) {
	proxy := httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme =
			req.URL.Host =
			req.URL.Path =
			req.Header =
			req.Host =
		},
	}

	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}

func StartInvocation(engine *gin.Engine) {
	engine.Any("/proxy", reverseProxyAPI)
}
