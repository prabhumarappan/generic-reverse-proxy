package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"reflect"
	"testing"
)


func TestCreateJSONFromError(t *testing.T) {
	err := errors.New("testing out")
	result := createJSONFromError(err)
	expectedOutput := map[string]string{
		"error": "testing out",
	}
	eq := reflect.DeepEqual(result, expectedOutput)
	if !eq {
		t.Error("testing out expected to be in map[string]string")
	}
}

func TestCheckRequestBody(t *testing.T) {
	inputRequestBody := RequestBodyFormat{
		ClientID:    "123",
		URL:         "www.google..com",
		Headers:     nil,
		RequestType: "",
		RequestBody: "",
	}
	err := checkRequestBody(inputRequestBody)

	if err.Error() != "empty request type provided" {
		t.Error("Issues with checking the request error type")
	}
}

func TestCheckRequestBodyValid(t *testing.T) {
	inputRequestBody := RequestBodyFormat{
		ClientID:    "123",
		URL:         "www.google..com",
		Headers:     nil,
		RequestType: "GET",
		RequestBody: "",
	}
	err := checkRequestBody(inputRequestBody)

	if err != nil {
		t.Error("Issues with checking the request error type")
	}
}

func TestCheckHTTPSRequest(t *testing.T) {
	inputRequestBody := RequestBodyFormat{
		ClientID:    "123",
		URL:         "www.google..com",
		Headers:     nil,
		RequestType: "GET",
		RequestBody: "",
	}
	err := checkHTTPSRequest(inputRequestBody)
	if err.Error() != "cannot process non-https URLs"{
		t.Error("Issues with checking the https request")
	}
}

func TestCheckHTTPSRequestValid(t *testing.T) {
	inputRequestBody := RequestBodyFormat{
		ClientID:    "123",
		URL:         "https://www.google..com",
		Headers:     nil,
		RequestType: "GET",
		RequestBody: "",
	}
	err := checkHTTPSRequest(inputRequestBody)
	if err != nil {
		t.Error("Issues with checking the https request")
	}
}

func TestGetRequestBody(t *testing.T) {
	data := map[string]string{
		"ClientID": "123",
		"URL": "https://www.duckduckgo.com",
	}
	dataBytes, _ := json.Marshal(data)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", bytes.NewBuffer(dataBytes))
	c, _ := gin.CreateTestContext(w)
	c.Request = r

	rb, _ := getRequestBody(c)
	if rb.ClientID != "123" || rb.URL != "https://www.duckduckgo.com" {
		t.Error("Error parsing the request body!")
	}
}