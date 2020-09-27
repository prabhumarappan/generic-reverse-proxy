package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func attachRequestBodyToContext(ctx *gin.Context, requestBody middleware.RequestBodyFormat) {
	ctx.Request.URL, _ = url.Parse(requestBody.URL)
	if requestBody.Headers != nil {
		for header, value := range requestBody.Headers {
			ctx.Request.Header.Set(header, value)
		}
	}
	ctx.Request.Method = requestBody.RequestType
	bodyReader := bytes.NewReader([]byte(requestBody.RequestBody))
	ctx.Request.Body = ioutil.NopCloser(bodyReader)
	ctx.Request.ContentLength = bodyReader.Size()
}

func reverseProxyAPI(ctx *gin.Context) {
	rb := ctx.MustGet("RequestBody")
	requestBody := rb.(middleware.RequestBodyFormat)
	attachRequestBodyToContext(ctx, requestBody)
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = ctx.Request.URL.Scheme
			req.URL.Host = ctx.Request.URL.Host
			req.URL.Path = ctx.Request.URL.Path

			req.Header = ctx.Request.Header
			req.Method = ctx.Request.Method
			req.Body = ctx.Request.Body
		},
	}

	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}

func StartInvocation(engine *gin.Engine) {
	engine.Any("/proxy", reverseProxyAPI)
}
