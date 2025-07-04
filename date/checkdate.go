package date

import (
	"errors"
	"graduation_project/db"
	"strconv"
	"strings"
	"time"
)

// проверка дат
func CheckDate(task *db.Task) error {
	now := time.Now()
	today := now.Truncate(24 * time.Hour)

	task.Date = strings.TrimSpace(task.Date)
	task.Title = strings.TrimSpace(task.Title)
	task.Repeat = strings.TrimSpace(task.Repeat)

	if task.Title == "" {
		return errors.New("title не может будет пустым")
	}

	if task.Date == "" {
		task.Date = now.Format(FormatDate)
	}

	if len(task.Date) != 8 {
		return errors.New("дата не должна содержать больше 8-ми символов")
	}

	for _, d := range task.Date {
		if d < '0' || d > '9' {
			return errors.New("некорректная дата")
		}
	}

	t, err := time.Parse(FormatDate, task.Date)
	if err != nil || t.Format(FormatDate) != task.Date {
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
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return err
			}
			task.Date = next
		}
	} else {
		if t.Before(today) {
			task.Date = now.Format(FormatDate)
		}
	}
	return nil

}
