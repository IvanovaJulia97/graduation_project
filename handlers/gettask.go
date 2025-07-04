package handlers

import (
	"graduation_project/db"
	"net/http"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func GetTasksHandlers(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.SortTask(50)
	if err != nil {
		WriteJSON(w, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, TasksResp{Tasks: tasks})

}
