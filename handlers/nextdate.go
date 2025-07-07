package handlers

import (
	"graduation_project/date"
	"net/http"
	"time"
)

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
	// w.Write([]byte(res))
	// WriteJSON(w, map[string]string{"next_date": res})
	_, err = w.Write([]byte(res))
	if err != nil {
		http.Error(w, "Ошибка записи ответа", http.StatusInternalServerError)
	}

}
