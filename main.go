package main

import (
	"fmt"
	"log"
	"log/slog"
	"photofiler/internal/ui"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	if err != nil {
		slog.Warn(fmt.Sprintf("Error loading .env file: %v", err))
	}
	ui.Run()
}
