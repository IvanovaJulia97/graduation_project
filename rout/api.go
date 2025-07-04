package rout

import (
	"graduation_project/handlers"
	"net/http"
)

func Init() {
	http.HandleFunc("/api/nextdate", handlers.NextDateHandler)
	http.HandleFunc("/api/task", handlers.TaskHandler)
	http.HandleFunc("/api/tasks", handlers.GetTasksHandlers)
	http.HandleFunc("/api/task/done", handlers.DoneTaskHandler)
}
