package main

import (
	"log"
	"os"

	"github.com/PolakoJambo/Final-project/pkg/db"
	"github.com/PolakoJambo/Final-project/pkg/server"
)

func main() {

	port := "7540"
	if envPort := os.Getenv("TODO_PORT"); envPort != "" {
		port = envPort
	}

	dbFile := "scheduler.db"
	if envDBFile := os.Getenv("TODO_DBFILE"); envDBFile != "" {
		dbFile = envDBFile
	}

	err := db.Init(dbFile)
	if err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}

	if err := server.StartServer(port); err != nil {
		log.Fatalf("server startup failed: %v", err)
	}
}
