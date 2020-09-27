package middleware

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type RequestBodyFormat struct {
	ClientID    string            `json:"ClientID"`
	URL         string            `json:"URL"`
	Headers     map[string]string `json:"Headers"`
	RequestType string            `json:"RequestType"`
	RequestBody string            `json:"RequestBody"`
}

func createJSONFromError(err error) map[string]string {
	errorMessage := err.Error()
	response := map[string]string{
		"error": errorMessage,
	}
	return response
}

func getRequestBody(ctx *gin.Context) (RequestBodyFormat, error) {
	var rb RequestBodyFormat
	requestBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return rb, errors.New("bad request body or no request body provided")
	}

	if len(requestBody) == 0 {
		return rb, errors.New("no request body provided")
	}

	err = json.Unmarshal(requestBody, &rb)
	if err != nil {
		return rb, errors.New("error with parsing the request body")
	}
	return rb, nil
}

func checkHTTPSRequest(rb RequestBodyFormat) error {
	URL, err := url.Parse(rb.URL)
	if err != nil {
		return err
	}
	if URL.Scheme != "https" {
		return errors.New("cannot process non-https URLs")
	}
	return nil
}

func checkRequestBody(rb RequestBodyFormat) error {
	if len(rb.URL) == 0 {
		return errors.New("empty URL path provided")
	}
	if len(rb.ClientID) == 0 {
		return errors.New("empty ClientId provided")
	}
	if (len(rb.RequestType)) == 0 {
		return errors.New("empty request type provided")
	}
	if rb.RequestType != "GET" && rb.RequestType != "POST" && rb.RequestType != "PUT" && rb.RequestType != "DELETE" {
		return errors.New("request type needs to be either GET or POST or PUT or DELETE")
	}
	return nil
}

func RequestParser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rb, err := getRequestBody(ctx)
		if err != nil {
			response := createJSONFromError(err)
			ctx.AbortWithStatusJSON(400, response)
			return
		}

		err = checkRequestBody(rb)
		if err != nil {
			response := createJSONFromError(err)
			ctx.AbortWithStatusJSON(400, response)
			return
		}

		err = checkHTTPSRequest(rb)
		if err != nil {
			response := createJSONFromError(err)
			ctx.AbortWithStatusJSON(400, response)
			return
		}

		ctx.Set("ClientId", rb.ClientID)
		ctx.Set("RequestBody", rb)
	}
}
