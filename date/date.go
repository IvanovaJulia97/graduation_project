package date

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const FormatDate = "20060102"

// функция возвращает true , если первая дата больше второй
func AfterNow(date, now time.Time) bool {
	date1 := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	date2 := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	return date1.After(date2)
}

// функция возвращает последний день месяца
func lastDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {

	//парсинг даты
	dateStart, err := time.Parse(FormatDate, dstart)
	if err != nil {
		return "", fmt.Errorf("ошибка при парсинге даты: %w", err)
	}

	if strings.TrimSpace(repeat) == "" {
		return "", errors.New("выполненная задача удалена, так как правило не указано")
	}

	//функция оставляет только дату
	normTime := func(t time.Time) time.Time {
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	}

	now = normTime(now)

	//парсим repeat
	strRepeat := strings.Fields(repeat)

	switch strRepeat[0] {
	case "d":
		if len(strRepeat) != 2 {
			return "", errors.New("ошибка, repeat должен состоять из двух элементов")
		}

		period, err := strconv.Atoi(strRepeat[1])
		if err != nil || period < 1 || period > 400 {
			return "", errors.New("ошибка, число должно быть от 1 до 400")
		}

		next := normTime(dateStart)
		now = normTime(now)

		//сравниваем, сколько прошло времени от тек даты и начальной даты
		daysBetween := int(now.Sub(next).Hours() / 24)
		if daysBetween < 0 {
			daysBetween = 0
		}

		periodPassed := daysBetween / period

		next = next.AddDate(0, 0, (periodPassed+1)*period)

		return next.Format(FormatDate), nil

	case "y":
		if len(strRepeat) != 1 {
			return "", errors.New("ошибка, y должен содержать 1 элемент")
		}

		origin := dateStart

		for {
			origin = origin.AddDate(1, 0, 0)

			last := lastDayOfMonth(origin)

			//вычисляем для каждого года последний день месяца
			var next time.Time
			if origin.Day() > last.Day() {
				next = last
			} else {
				next = time.Date(origin.Year(), origin.Month(), origin.Day(), 0, 0, 0, 0, time.UTC)
			}

			if AfterNow(normTime(next), now) {
				return next.Format(FormatDate), nil
			}
		}

	default:
		return "", errors.New("переданный формат даты не поддерживается")
	}
}
