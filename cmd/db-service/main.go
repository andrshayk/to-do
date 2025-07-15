package main

import (
	"log"
	"net/http"

	"github.com/andrdog/nilchan/internal/db/handler"
	"github.com/andrdog/nilchan/internal/db/storage"
)

func main() {
	if err := storage.InitPostgres(); err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer storage.DB.Close()

	http.HandleFunc("/list", handler.ListTasks)

	log.Println("db-service started on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
