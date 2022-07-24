package foodb

import (
	"database/sql"

	"wagie.com/wageslavery/resources"
	"wagie.com/wageslavery/viewmodels"
)

type FooRepository struct{}

func (f *FooRepository) Get() ([]viewmodels.FooGet, error) {
	var res []viewmodels.FooGet
	db := resources.GetSQLite()
	query := `
		SELECT 
			id,
			key,
			value
		FROM
			foo
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {

		var id sql.NullInt64
		var key, value sql.NullString

		if err := rows.Scan(
			&id,
			&key,
			&value,
		); err != nil {
			return nil, err
		}

		res = append(res, viewmodels.FooGet{
			Id:    id.Int64,
			Key:   key.String,
			Value: value.String,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

func (f *FooRepository) Create(req viewmodels.FooGet) error {
	db := resources.GetSQLite()

	_, err := db.Exec(
		`INSERT INTO foo (key,value) VALUES (?,?);`,
		req.Key,
		req.Value,
	)
	if err != nil {
		return err
	}
	return nil
}
