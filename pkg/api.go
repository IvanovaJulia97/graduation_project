package pkg

import "net/http"

func Init() {
	http.HandleFunc("/api/nextdate", nextDateHandler)
}
