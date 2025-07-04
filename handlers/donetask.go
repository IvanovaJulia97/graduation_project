package handlers

import (
	"graduation_project/date"
	"graduation_project/db"
	"net/http"
	"time"
)

func DoneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		WriteJSON(w, map[string]string{"error": "не указан id задачи"})
		return
	}

	task, err := db.GetTasks(id)
	if err != nil {
		WriteJSON(w, map[string]string{"error": "задачи с таким id не найдено"})
		return
	}

	if task.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			WriteJSON(w, map[string]string{"error": "ошибка правила повтора"})
			return
		}
		WriteJSON(w, map[string]string{})
		return
	}

	nextDate, err := date.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		WriteJSON(w, map[string]string{"error": "некорректная дата"})
		return
	}

	if err := db.UpdateDate(nextDate, id); err != nil {
		WriteJSON(w, map[string]string{"error": "ошибка обновления даты"})
		return
	}

	WriteJSON(w, map[string]string{})

}
