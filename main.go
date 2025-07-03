package main

import (
	"graduation_project/db"
	"graduation_project/rout"
	"graduation_project/server"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
	}

	log.Println("Успешное подключение к БД")

	rout.Init()

	server.MyServer()
}
