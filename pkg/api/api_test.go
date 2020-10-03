package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/prabhumarappan/freshworks-hiring/pkg/middleware"
	"io/ioutil"
	"net/http/httptest"
	"net/url"
	"testing"
)

func setup() (*gin.Context, middleware.RequestBodyFormat) {
	rb := middleware.RequestBodyFormat{
		ClientID:    "123",
		URL:         "https://www.duckduckgo.com",
		Headers:     map[string]string{
			"Name": "Prabhu",
		},
		RequestType: "PUT",
		RequestBody: "Hello World",
	}
	rbBytes, _ := json.Marshal(rb)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/proxy", ioutil.NopCloser(bytes.NewBuffer(rbBytes)))
	c, _ := gin.CreateTestContext(w)
	c.Request = r

	return c, rb
}

func TestAttachRequestBodyToContext(t *testing.T) {
	c, rb := setup()
	attachRequestBodyToContext(c,  rb)

	if c.Request.Method != rb.RequestType {
		t.Error("Request Method was not attached properly!")
	}

	if c.Request.Header.Get("Name") != rb.Headers["Name"] {
		t.Error("Request Header was not attached properly!")
	}

	parsedURL, _ := url.Parse(rb.URL)
	if c.Request.URL.Host != parsedURL.Host {
		t.Error("Request URL has not been set properly!")
	}

	if c.Request.ContentLength != int64(len(rb.RequestBody)) {
		t.Error("Request content length has not been set properly")
	}

	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	bodyString := string(bodyBytes)
	if bodyString != rb.RequestBody {
		t.Error("Request body was not set properly")
	}
}
