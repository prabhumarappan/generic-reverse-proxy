package middleware

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

type RequestBodyFormat struct {
	ClientID string `json:"ClientID"`
	URL string `json:"URL"`
	Headers map[string]interface{} `json:"Headers"`
	RequestType string `json:"RequestType"`
	RequestBody string `json:"RequestBody"`
}

func checkRequestBody(ctx *gin.Context) error {
	requestBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error with reading Request Body")
		return err
	}

}

func checkHTTPSRequest(ctx *gin.Context) error {

}

func RequestParser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := checkRequestBody(ctx)
		if err != nil {
			ctx.AbortWithError(400, err)
		}
		err = checkHTTPSRequest(ctx)
		if err != nil {
			ctx.AbortWithError(400, err)
		}
	}
}