package main

import (
	"graduation_project/db"
	"graduation_project/rout"
	"graduation_project/server"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	// err1 := godotenv.Load()
	// if err1 != nil {
	// 	log.Println("Ошибка загрузки .env файда")
	// }

	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Ошибка в подключении к БД: %v", err)
	}

	log.Println("Успешное подключение к БД")

	rout.Init()

	server.MyServer()
}
