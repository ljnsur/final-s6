package main

import (
	"log"
	"os"

	"s6-final/internal/server"
)

func main() {
	file, err := os.OpenFile("full.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	InfoLogger := log.New(file, "INFO: ", log.LstdFlags|log.Lshortfile)

	newServer := server.New(InfoLogger)

	if err := newServer.Start(); err != nil {
		InfoLogger.Fatal(err)
	}

}
