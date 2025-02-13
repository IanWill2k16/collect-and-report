package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/IanWill2k16/collect-and-report/internal/database"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	dbQueries   *database.Queries
	envPlatform string
}

func main() {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	envPlatformVar := os.Getenv("PLATFORM")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("error connect to database: %v", err)
	}

	apiCfg := &apiConfig{
		dbQueries:   database.New(db),
		envPlatform: envPlatformVar,
	}

	const port = "8080"
	mux := http.NewServeMux()
	srv := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	mux.HandleFunc("POST /api/users", apiCfg.createUser)

	log.Fatal(srv.ListenAndServe())
}
