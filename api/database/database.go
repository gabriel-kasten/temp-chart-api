package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Result struct {
	ID           int     `json:"id"`
	DateMeasured string  `json:"date_measured"`
	Temperature  float64 `json:"temperature"`
}

func DatabaseCon() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

func QueryTemperatureData(db *sql.DB) ([]Result, error) {
	var results []Result

	rows, err := db.Query("SELECT id, date_measured, temperature FROM temperature_data")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r Result
		err = rows.Scan(&r.ID, &r.DateMeasured, &r.Temperature)
		if err != nil {
			return nil, err
		}
		results = append(results, r)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, nil

}
