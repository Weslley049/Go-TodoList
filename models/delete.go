package models

import "github.com/Weslley049/todo-list/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, nil
	}

	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
