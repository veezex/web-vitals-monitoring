package main

import (
	"fmt"
	"github.com/veezex/web-vitals-monitoring/server/internal/pkg/config"
	"github.com/veezex/web-vitals-monitoring/server/internal/pkg/db"
	"github.com/veezex/web-vitals-monitoring/server/internal/pkg/server"
	"log"
)

const DB_PATH = "./metrics.db"

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatalln("Failed to load config:", err)
	}

	newDb, dbErr := db.New(DB_PATH)
	if dbErr != nil {
		log.Fatalln("Failed to create database:", dbErr)
	}
	defer newDb.Close()

	serverErr := server.RunServer(config, newDb)
	if serverErr != nil {
		fmt.Printf("Error starting server: %s\n", serverErr)
	}
}
