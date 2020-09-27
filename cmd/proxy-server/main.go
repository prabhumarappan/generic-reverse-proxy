package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prabhumarappan/freshworks-hiring/pkg/api"
	"github.com/prabhumarappan/freshworks-hiring/pkg/middleware"
	"log"
	"net/http"
	"time"
)

func createServer() {
	server := gin.Default()
	server.Use(gin.Logger())
	server.Use(middleware.RequestParser())

	api.StartInvocation(server)

	srv := &http.Server{
		Addr:              ":8083",
		Handler:           server,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func main()  {
	createServer()
}