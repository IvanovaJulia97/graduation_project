package handlers

import (
	"encoding/json"
	"graduation_project/date"
	"graduation_project/db"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "ошибка кодирования JSON", http.StatusInternalServerError)
	}
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		AddTaskHandler(w, r)
	case http.MethodGet:
		id := r.URL.Query().Get("id")
		if id == "" {
			WriteJSON(w, map[string]string{"error": "не передан id задачи"})
			return
		}
		task, err := db.GetTasks(id)
		if err != nil {
			WriteJSON(w, map[string]string{"error": "ошибка получения задачи"})
			return
		}
		WriteJSON(w, task)

	case http.MethodPut:
		var task db.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			WriteJSON(w, map[string]string{"error": "ошибка преобразования в JSON"})
			return
		}
		if err := date.CheckDate(&task); err != nil {
			WriteJSON(w, map[string]string{"error": "некорректная дата"})
			return
		}
		err := db.UpdateTask(&task)
		if err != nil {
			WriteJSON(w, map[string]string{"error": "ошибка при обновлении задачи"})
			return
		}
		WriteJSON(w, map[string]string{})
	case http.MethodDelete:
		id := r.URL.Query().Get("id")
		if id == "" {
			WriteJSON(w, map[string]string{"error": "не указан id задачи"})
			return
		}
		err := db.DeleteTask(id)
		if err != nil {
			WriteJSON(w, map[string]string{"error": "ошибка при удалении задачи"})
			return
		}
		WriteJSON(w, map[string]string{})

	default:
		w.WriteHeader(http.StatusBadRequest)
		//http.Error(w, "Данные запрос не поддерживается", http.StatusBadRequest)
		WriteJSON(w, map[string]string{"error": "данный запрос не поддерживается"})
	}
}
