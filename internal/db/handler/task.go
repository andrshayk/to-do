package handler

import (
	"net/http"

	"github.com/andrdog/nilchan/internal/db/storage"
	"github.com/gin-gonic/gin"
)

// Task представляет задачу чек-листа.
// @Description Структура задачи
type Task struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

func AutoMigrate() {
	storage.DB.AutoMigrate(&Task{})
}

// ListTasks возвращает список всех задач.
// @Summary Получить список задач
// @Description Возвращает все задачи из базы данных
// @Tags tasks
// @Produce json
// @Success 200 {array} Task
// @Router /list [get]
func ListTasks(c *gin.Context) {
	var tasks []Task
	storage.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}
