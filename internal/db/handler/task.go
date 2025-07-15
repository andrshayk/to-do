package handler

import (
	"net/http"

	"github.com/andrdog/nilchan/internal/db/storage"
	"github.com/gin-gonic/gin"
)

type Task struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

func AutoMigrate() {
	storage.DB.AutoMigrate(&Task{})
}

func ListTasks(c *gin.Context) {
	var tasks []Task
	storage.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}
