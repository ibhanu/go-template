package config

import (
	"log"
	"sync"

	"web-server/prisma/db"
)

// GetPrismaClient returns a singleton instance of PrismaClient.
var GetPrismaClient = (func() func() *db.PrismaClient {
	var (
		once   sync.Once
		client *db.PrismaClient
	)

	return func() *db.PrismaClient {
		once.Do(func() {
			client = db.NewClient()
			if err := client.Connect(); err != nil {
				log.Fatalf("Could not connect to database: %v", err)
			}
		})
		return client
	}
})()

// DisconnectDB closes the database connection.
func DisconnectDB() {
	if client := GetPrismaClient(); client != nil {
		if err := client.Disconnect(); err != nil {
			log.Printf("Error disconnecting from database: %v", err)
		}
	}
}
