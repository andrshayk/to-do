package main

import (
	"log"
	"net/http"

	"github.com/andrdog/nilchan/internal/db/handler"
)

func main() {
	http.HandleFunc("/list", handler.ListTasks)

	log.Println("db-service started on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
