package main

import (
	"log"

	"github.com/andrdog/nilchan/internal/db/handler"
	"github.com/andrdog/nilchan/internal/db/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := storage.InitGorm(); err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	handler.AutoMigrate()

	r := gin.Default()
	r.GET("/list", handler.ListTasks)

	log.Println("db-service started on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
