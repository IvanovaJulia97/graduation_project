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

// d <число> - максимальное число 400. где число это дни
// y - ежегодное выполнение задачи
func NextDate(now time.Time, dstart string, repeat string) (string, error) {

	//парсинг даты
	dateStart, err := time.Parse(FormatDate, dstart)
	if err != nil {
		return "", fmt.Errorf("ошибка при парсинге даты: %w", err)
	}

	//проверка на пустую строку repeat
	if strings.TrimSpace(repeat) == "" {
		return "", errors.New("выполненная задача удалена, так как правило не указано")
	}

	//для сравнения дат по дню
	normTime := func(t time.Time) time.Time {
		return time.Date(
			t.Year(),
			t.Month(),
			t.Day(),
			0, 0, 0, 0,
			time.UTC)
	}

	now = normTime(now)

	//парсим repeat
	strRepeat := strings.Fields(repeat)

	switch strRepeat[0] {
	case "d":
		if len(strRepeat) != 2 {
			return "", errors.New("ошибка, правило повторения <d> должно состоять из двух элементов")
		}

		period, err := strconv.Atoi(strRepeat[1])
		if err != nil || period < 1 || period > 400 {
			return "", errors.New("ошибка, у правила повторения <d> число должно быть от 1 до 400")
		}

		//приведение к дате, без времени
		next := normTime(dateStart)
		now = normTime(now)

		//вычисление дней между двумя датами
		daysBetween := int(now.Sub(next).Hours() / 24)
		if daysBetween < 0 {
			daysBetween = 0
		}

		periodPassed := daysBetween / period

		next = next.AddDate(0, 0, (periodPassed+1)*period)

		return next.Format(FormatDate), nil

	case "y":
		if len(strRepeat) != 1 {
			return "", errors.New("ошибка, <y> должен содержать 1 элемент")
		}

		//старт вычисления повторов
		origin := dateStart

		for {
			//сдвиг на год вперед
			origin = origin.AddDate(1, 0, 0)

			//получаем последний день месяца
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

	case "w":
		dayStr := strings.Split(strRepeat[1], ",")
		days := make([]int, 0, len(dayStr))
		for _, d := range dayStr {
			dayNum, err := strconv.Atoi(d)
			if err != nil || dayNum < 1 || dayNum > 7 {
				return "", errors.New("дни недель должны быть с числами от 1 до 7")
			}
			days = append(days, dayNum)
		}

		weekday := int(now.Weekday())
		if weekday == 0 {
			weekday = 7
		}

		minDays := 8
		for _, d := range days {
			diff := d - weekday
			if diff < 0 {
				diff += 7
			}
			if diff < minDays {
				minDays = diff
			}
		}

		next := now.AddDate(0, 0, minDays)
		return next.Format(FormatDate), nil
	case "m":
		if len(strRepeat) < 2 || len(strRepeat) > 3 {
			return "", errors.New("некорректно указано правило повторения <m>")
		}
		dayStr := strings.Split(strRepeat[1], ",")
		days := make([]int, 0, len(dayStr))
		for _, d := range dayStr {
			dayNum, err := strconv.Atoi(d)
			if err != nil || (dayNum < -2 || dayNum == 0 || dayNum > 31) {
				return "", errors.New("дни месяца должны быть от 1 до 31 или -1 или -2")
			}
			days = append(days, dayNum)
		}

		var months []int
		if len(strRepeat) == 3 {
			monthSTr := strings.Split(strRepeat[2], ",")
			months = make([]int, 0, len(monthSTr))
			for _, m := range monthSTr {
				monthNum, err := strconv.Atoi(m)
				if err != nil || monthNum < 1 || monthNum > 12 {
					return "", errors.New("количество месяцев должно быть от 1 до 12")
				}
				months = append(months, monthNum)
			}
		}

		for i := 0; i < 400; i++ {
			cur := now.AddDate(0, 0, i)
			month := int(cur.Month())
			day := cur.Day()

			if len(months) > 0 {
				foundMonth := false
				for _, m := range months {
					if m == month {
						foundMonth = true
					}
				}
				if !foundMonth {
					continue
				}
			}
			lastDay := lastDayOfMonth(cur).Day()
			secondLastDay := lastDay - 1

			for _, d := range days {
				switch {
				case d > 0:
					if day == d {
						return cur.Format(FormatDate), nil
					}
				case d == -1:
					if day == lastDay {
						return cur.Format(FormatDate), nil
					}
				case d == -2:
					if day == secondLastDay {
						return cur.Format(FormatDate), nil
					}
				}
			}
		}
		return "", errors.New("ненайдена дата для повторения")

	default:
		return "", errors.New("переданный формат даты не поддерживается")
	}
}
