package model

import (
	"database/sql"
	"fmt"
)

func connect(dbname string) (*sql.DB, error) {
	database, _ := sql.Open("sqlite3", fmt.Sprintf("./db/%s.db", dbname))

	statement, err := database.Prepare(CREATE_SUBMISSION_TABLE)
	if err != nil {
		return database, err
	}
	statement.Exec()

	statement, err = database.Prepare(CREATE_SUBMISSION_INDEX1)
	if err != nil {
		return database, err
	}
	statement.Exec()

	statement, err = database.Prepare(CREATE_SUBMISSION_INDEX2)
	if err != nil {
		return database, err
	}
	statement.Exec()

	return database, nil
}
