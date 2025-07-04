package handlers

import (
	"encoding/json"
	"graduation_project/date"
	"graduation_project/db"
	"net/http"
	"strconv"
	"strings"
)

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {

	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		http.Error(w, `{"error":"ожидался application/json"}`, http.StatusBadRequest)
		return
	}

	var task db.Task

	//десериализация JSON
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		WriteJSON(w, map[string]string{"error": "в преобразовании JSON"})
		return
	}

	//проверка дат
	if err := date.CheckDate(&task); err != nil {
		WriteJSON(w, map[string]string{"error": err.Error()})
		return
	}

	//проверка добавление задачи
	id, err := db.AddTask(db.DB, &task)
	if err != nil {
		WriteJSON(w, map[string]string{"error": err.Error()})
		return
	}

	//fmt.Printf("DEBUG: Получена задача: %+v\n", task)

	WriteJSON(w, map[string]string{"id": strconv.FormatInt(id, 10)})

}
