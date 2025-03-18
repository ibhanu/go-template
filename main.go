package main

import (
"log"
"os"
_ "web-server/docs" // swagger docs
"web-server/internal/infrastructure/server"

_ "github.com/swaggo/files"
_ "github.com/swaggo/gin-swagger"
)

// @title Web Server API
// @version 1.0
// @description A REST API server written in Go using gin-gonic framework.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	srv := server.NewServer()
	log.Fatal(srv.Start(":" + port))
}
