package pkg

import (
	"graduation_project/date"
	"net/http"
)

func Init() {
	http.HandleFunc("/api/nextdate", date.NextDateHandler)
}
