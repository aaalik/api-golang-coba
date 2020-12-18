package model

import (
	"github.com/aaalik/coba-golang/bootstrap"
	"github.com/aaalik/coba-golang/helper"
	"github.com/aaalik/coba-golang/structs"
)

func GetSingleItem(id int) (structs.Item, error) {
	sql := `
		SELECT
			id,
			name,
			price
		FROM
			item
		WHERE
			id = ?
	`

	rows := bootstrap.DB.QueryRow(sql, id)

	result := structs.Item{}

	err := rows.Scan(
		&result.Id,
		&result.Name,
		&result.Price,
	)

	return result, err
}

func GetItem() []structs.Item {
	sql := `
		SELECT
			id,
			name,
			price
		FROM
			item
	`

	rows, err := bootstrap.DB.Query(sql)
	if err != nil {
		helper.Log.Error(err)
	}

	results := []structs.Item{}

	for rows.Next() {
		row := structs.Item{}

		err := rows.Scan(
			&row.Id,
			&row.Name,
			&row.Price,
		)

		if err != nil {
			helper.Log.Error(err)
		}

		results = append(results, row)
	}

	return results
}

func SaveItem(name string, price int) (bool, error) {
	sql := `
		INSERT INTO item(name, price) VALUES(?,?)
	`

	rows, err := bootstrap.DB.Prepare(sql)

	if err != nil {
		helper.Log.Error(err)
		return false, err
	}

	rows.Exec(name, price)

	return true, err
}

func UpdateItem(id int, name string, price int) (bool, error) {
	sql := `
		UPDATE item SET name = ? , price = ? WHERE id = ?
	`

	rows, err := bootstrap.DB.Prepare(sql)

	if err != nil {
		helper.Log.Error(err)
		return false, err
	}

	rows.Exec(name, price, id)

	return true, err
}

func DeleteItem(id int) (bool, error) {
	sql := `
		DELETE FROM item WHERE id = ?
	`

	rows, err := bootstrap.DB.Prepare(sql)

	if err != nil {
		helper.Log.Error(err)
		return false, err
	}

	rows.Exec(id)

	return true, err
}
