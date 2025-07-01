package main

import (
	"graduation_project/db"
	"graduation_project/pkg"
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

	pkg.Init()

	server.MyServer()
}
