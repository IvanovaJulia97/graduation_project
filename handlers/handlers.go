package handlers

import (
	"encoding/json"
	"graduation_project/date"
	"graduation_project/db"
	"net/http"
	"strconv"
	"time"
)

func WriteJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
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

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		AddTaskHandler(w, r)
	default:
		http.Error(w, "Данные запрос не поддерживается", http.StatusBadRequest)

	}
}

func NextDateHandler(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	startDate := r.FormValue("date")
	repeat := r.FormValue("repeat")

	var nowTime time.Time
	var err error

	if now == "" {
		nowTime = time.Now()
	} else {
		nowTime, err = time.Parse(date.FormatDate, now)
		if err != nil {
			http.Error(w, "Неверный формат даты", http.StatusBadRequest)
			return
		}
	}

	res, err := date.NextDate(nowTime, startDate, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))

}
