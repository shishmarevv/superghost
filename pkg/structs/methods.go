package structs

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func (s Sequence) Update() {
	db := open()
	checkDB(db)
	defer shut(db)

	rows, err := db.Query(`
		SELECT weight
		FROM weights
		WHERE sequence = ?
		`, s.Text)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&s.Weight)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func open() *sql.DB {
	db, err := sql.Open("sqlite3", "superghost/pkg/mem/mem.db")
	checkErr(err)

	log.Println("Database opened")

	return db
}

func shut(db *sql.DB) {
	err := db.Close()
	checkErr(err)

	log.Println("Database closed")
}

func checkDB(db *sql.DB) bool {
	err := db.Ping()
	checkErr(err)

	return true
}
