package handlers

import (
	"encoding/json"
	"graduation_project/date"
	"net/http"
	"time"
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
	default:
		w.WriteHeader(http.StatusBadRequest)
		//http.Error(w, "Данные запрос не поддерживается", http.StatusBadRequest)
		WriteJSON(w, map[string]string{"error": "Данный запрос не поддерживается"})
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
			w.WriteHeader(http.StatusBadRequest)
			WriteJSON(w, map[string]string{"error": "Неверный формат даты"})
			return
		}
	}

	res, err := date.NextDate(nowTime, startDate, repeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		WriteJSON(w, map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	//w.Write([]byte(res))
	WriteJSON(w, map[string]string{"next_date": res})

}
