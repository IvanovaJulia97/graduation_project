package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"graduation_project/date"
	"graduation_project/db"
	"graduation_project/tasks"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func WriteJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)
}

func CheckDate(task *tasks.Task) error {
	now := time.Now()
	today := now.Truncate(24 * time.Hour)

	task.Date = strings.TrimSpace(task.Date)
	task.Title = strings.TrimSpace(task.Title)
	task.Repeat = strings.TrimSpace(task.Repeat)

	if task.Title == "" {
		return errors.New("title не может будет пустым")
	}

	if task.Date == "" {
		task.Date = now.Format(date.FormatDate)
	}

	if len(task.Date) != 8 {
		return errors.New("дата не должна содержать больше 8-ми символов")
	}

	for _, d := range task.Date {
		if d < '0' || d > '9' {
			return errors.New("некорректная дата")
		}
	}

	t, err := time.Parse(date.FormatDate, task.Date)
	if err != nil || t.Format(date.FormatDate) != task.Date {
		return errors.New("некорректный формат даты")
	}

	if task.Repeat != "" {
		p := strings.Fields(task.Repeat)

		if len(p) == 0 || len(p) > 2 {
			return errors.New("повтор имеет неверную длину")
		}

		if p[0] != "d" && p[0] != "y" {
			return errors.New("повтор должен содержать только d и y")
		}

		if p[0] == "y" && len(p) != 1 {
			return errors.New("повтор y должен содержать 1 элемент")
		}

		if p[0] == "d" {
			if len(p) != 2 {
				return errors.New("повтор должен содержать число")
			}
			n, err := strconv.Atoi(p[1])
			if err != nil || n < 1 || n > 400 {
				return errors.New("некорректный повтор")
			}
		}

		if t.Before(today) {
			next, err := date.NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return err
			}
			task.Date = next
		}
	} else {
		if t.Before(today) {
			task.Date = now.Format(date.FormatDate)
		}
	}
	return nil

}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task tasks.Task

	//десериализация JSON
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		WriteJSON(w, map[string]string{"error": "в преобразовании JSON"})
		return
	}

	//проверка дат
	if err := CheckDate(&task); err != nil {
		WriteJSON(w, map[string]string{"error": err.Error()})
		return
	}

	//проверка добавление задачи
	id, err := tasks.AddTask(db.DB, &task)
	if err != nil {
		WriteJSON(w, map[string]string{"error": err.Error()})
		return
	}

	fmt.Printf("DEBUG: Получена задача: %+v\n", task)

	WriteJSON(w, map[string]string{"id": strconv.FormatInt(id, 10)})

}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		AddTaskHandler(w, r)
	default:
		http.Error(w, "Данные запрос не поддерживается", http.StatusBadRequest)

	}
}
