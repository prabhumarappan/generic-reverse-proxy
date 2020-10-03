package api

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	limiter "github.com/julianshen/gin-limiter"
	"github.com/prabhumarappan/freshworks-hiring/pkg/middleware"
)

const PerHourRateLimit = 50

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
	lm := limiter.NewRateLimiter(time.Minute, PerHourRateLimit, func(ctx *gin.Context) (string, error) {
		key := ctx.GetString("ClientId")
		if key != "" {
			return key, nil
		}
		return "", errors.New("ClientId is missing!")
	})

	engine.Any("/proxy", lm.Middleware(), reverseProxyAPI)
}
