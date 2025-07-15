package handler

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []Task{
		{ID: 1, Title: "Пример задачи", Content: "Описание задачи", Done: false},
		{ID: 2, Title: "Вторая задача", Content: "Что-то сделать", Done: true},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
