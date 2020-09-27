package middleware

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/url"
)

type RequestBodyFormat struct {
	ClientID string `json:"ClientID"`
	URL string `json:"URL"`
	Headers map[string]interface{} `json:"Headers"`
	RequestType string `json:"RequestType"`
	RequestBody string `json:"RequestBody"`
}

func getRequestBody(ctx *gin.Context) (RequestBodyFormat, error) {
	var rb RequestBodyFormat
	requestBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error with reading Request Body")
		return rb, err
	}
	err = json.Unmarshal(requestBody, &rb)
	if err != nil {
		log.Println("Error with parsing request body")
		return rb, err
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
	return nil
}

func RequestParser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rb, err := getRequestBody(ctx)
		if err != nil {
			ctx.AbortWithError(400, err)
		}

		err = checkRequestBody(rb)
		if err != nil {
			ctx.AbortWithError(400, err)
		}

		err = checkHTTPSRequest(rb)
		if err != nil {
			ctx.AbortWithError(400, err)
		}
	}
}