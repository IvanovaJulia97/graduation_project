package date

import (
	"net/http"
	"time"
)

//const formatDate = "20060102"

func NextDateHandler(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	startDate := r.FormValue("date")
	repeat := r.FormValue("repeat")

	var nowTime time.Time
	var err error

	if now == "" {
		nowTime = time.Now()
	} else {
		nowTime, err = time.Parse(FormatDate, now)
		if err != nil {
			http.Error(w, "Неверный формат даты", http.StatusBadRequest)
			return
		}
	}

	res, err := NextDate(nowTime, startDate, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))

}
