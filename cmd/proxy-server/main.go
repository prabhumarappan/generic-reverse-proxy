package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"github.com/prabhumarappan/freshworks-hiring/pkg/api"
)

func createServer() {
	server := gin.Default()
	api.StartInvocation(&server)
	log.Fatal(server.Run(":8083"))
}

func main()  {
	createServer()
}