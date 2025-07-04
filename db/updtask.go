package db

import "fmt"

func UpdateTask(task *Task) error {
	req := `
		UPDATE scheduler
		SET date = ?, title = ?, comment = ?, repeat = ?
		WHERE id = ?`

	res, err := DB.Exec(req, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("некорректный id задачи для обновления")
	}

	return nil

}

func UpdateDate(next string, id string) error {
	req := `UPDATE scheduler SET date = ? WHERE id = ?`

	res, err := DB.Exec(req, next, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("задача с id = %s не найдена, обновление невозможно", id)
	}
	return nil

}
