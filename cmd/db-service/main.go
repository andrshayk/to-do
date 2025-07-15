package main

import (
	"log"

	_ "github.com/andrdog/nilchan/docs"
	"github.com/andrdog/nilchan/internal/db/handler"
	"github.com/andrdog/nilchan/internal/db/storage"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	if err := storage.InitGorm(); err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	handler.AutoMigrate()

	r := gin.Default()
	r.GET("/list", handler.ListTasks)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("db-service started on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
