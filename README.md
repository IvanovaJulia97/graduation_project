# Файлы для итогового задания

В директории `tests` находятся тесты для проверки API, которое должно быть реализовано в веб-сервере.

Директория `web` содержит файлы фронтенда.

# Описание проекта Планировщик задач

В рамках проекта было реализовано добавление, удаление и редактирование задач, а так же возможность выбора правил повтора задач

# Комнда для запуска кода локально
go run main.go

# Хост 
http://localhost:7540/

# Команда для запуска тестов
go test ./tests

# Команда для запуска проекта через докер
docker build -t scheduler-app .

docker run -d --name scheduler-container \
  -p 7540:7540 \
  -v $(pwd)/scheduler.db:/data/scheduler.db \
  -e TODO_PORT=7540 \
  -e TODO_DBFILE=/data/scheduler.db \
  -e TODO_PASSWORD=supersecret \
  scheduler-app

