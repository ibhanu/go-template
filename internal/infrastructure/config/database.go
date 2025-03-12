package config

import (
	"log"
	"web-server/prisma/db"
)

var prismaClient *db.PrismaClient

// GetPrismaClient returns a singleton instance of PrismaClient
func GetPrismaClient() *db.PrismaClient {
	if prismaClient == nil {
		client := db.NewClient()
		if err := client.Connect(); err != nil {
			log.Fatalf("Could not connect to database: %v", err)
		}

		prismaClient = client
	}

	return prismaClient
}

// DisconnectDB closes the database connection
func DisconnectDB() {
	if prismaClient != nil {
		if err := prismaClient.Disconnect(); err != nil {
			log.Printf("Error disconnecting from database: %v", err)
		}
	}
}
