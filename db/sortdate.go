package db

import "fmt"

func SortTask(limit int) ([]*Task, error) {
	str, err := DB.Query(`
		SELECT id, date, title, comment, repeat
		FROM scheduler
		ORDER BY CAST(date AS INT) ASC
		LIMIT ?
		`, limit)

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	defer str.Close()

	var tasks []*Task

	for str.Next() {
		var t Task
		err := str.Scan(&t.ID,
			&t.Date,
			&t.Title,
			&t.Comment,
			&t.Repeat)
		if err != nil {
			return nil, fmt.Errorf("ошибка получения задачи: %w", err)
		}
		//fmt.Println(">>", t.Date)
		tasks = append(tasks, &t)

	}

	if tasks == nil {
		tasks = []*Task{}
	}

	return tasks, nil

}
