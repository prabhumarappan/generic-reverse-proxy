package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prabhumarappan/freshworks-hiring/pkg/api"
	"github.com/prabhumarappan/freshworks-hiring/pkg/middleware"
	"log"
	"net/http"
	"time"
)

const TimeOutConstant = 5

func createServer() {
	server := gin.Default()
	server.Use(gin.Logger())
	server.Use(middleware.RequestParser())

	api.StartInvocation(server)

	srv := &http.Server{
		Addr:              ":8083",
		Handler:           server,
		ReadTimeout:       TimeOutConstant * time.Second,
		WriteTimeout:      TimeOutConstant * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func main()  {
	createServer()
}