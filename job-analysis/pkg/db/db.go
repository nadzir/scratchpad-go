package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "password"
	dbName   = "analysis"
)

func dbConn() (db *sql.DB) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	db.SetMaxOpenConns(20) // Sane default
	db.SetMaxIdleConns(0)
	db.SetConnMaxLifetime(time.Nanosecond)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// Select: Selecting from job table
func Select(query string, channel chan<- *sql.Rows) {
	db := dbConn()
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		// fmt.Println(rows)
		channel <- rows
		// err := rows.Scan(&jobTitle, &count)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// json.Set(jobTitle, count)
	}
	// return json
}
