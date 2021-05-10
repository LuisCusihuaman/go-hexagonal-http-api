package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const httpAddr = ":8000"

func main() {
	fmt.Println("Server running on", httpAddr)

	srv := gin.New()
	srv.GET("/health", healthHandler)

	log.Fatal(srv.Run(httpAddr))
}

func healthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "everything is ok 🚀")
}
