package main

import (
"log"
"web-server/internal/infrastructure/server"
)

func main() {
srv := server.NewServer()
log.Fatal(srv.Start(":8080"))
}