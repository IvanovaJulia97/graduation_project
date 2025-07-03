package rout

import (
	"graduation_project/date"
	"graduation_project/handlers"
	"net/http"
)

func Init() {
	http.HandleFunc("/api/nextdate", date.NextDateHandler)
	http.HandleFunc("/api/task", handlers.TaskHandler)
}
