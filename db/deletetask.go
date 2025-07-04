package db

import "fmt"

func DeleteTask(id string) error {
	req := `DELETE FROM scheduler WHERE id = ?`
	res, err := DB.Exec(req, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("задача с id = %s не найдена", id)
	}

	return nil

}
