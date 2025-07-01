package pkg

import (
	"graduation_project/date"
	"net/http"
	"time"
)

const formatDate = "20060102"

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	startDate := r.FormValue("date")
	repeat := r.FormValue("repeat")

	var nowTime time.Time
	var err error

	if now == "" {
		nowTime = time.Now()
	} else {
		nowTime, err = time.Parse(formatDate, now)
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
