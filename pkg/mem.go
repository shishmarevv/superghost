package superghost

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Open() *sql.DB {
	db, err := sql.Open("sqlite3", "superghost/memory/mem.db")
	checkErr(err)

	log.Println("Database opened")

	return db
}

func Shut(db *sql.DB) {
	err := db.Close()
	checkErr(err)

	log.Println("Database closed")
}

func CheckDB(db *sql.DB) bool {
	err := db.Ping()
	checkErr(err)

	return true
}

func AddWord(db *sql.DB, word string) {
	CheckDB(db)

	_, err := db.Exec(`
		INSERT INTO words(word) 
		VALUES(?)
		`, word)
	checkErr(err)

	log.Println("Added word:", word)
}

func IsInWord(db *sql.DB, include string) bool {
	CheckDB(db)

	include = "%" + include + "%"
	rows, err := db.Query(`
		SELECT word 
		FROM words 
		WHERE word LIKE ?
		`, include)
	checkErr(err)

	defer rows.Close()

	if rows.Next() {
		return true
	}
	return false
}

func IsWord(db *sql.DB, word string) bool {
	CheckDB(db)

	rows, err := db.Query(`
		SELECT word 
		FROM words 
		WHERE word == ?
		`, word)
	checkErr(err)

	defer rows.Close()

	if rows.Next() {
		return true
	}
	return false
}

func AddSequence(db *sql.DB, sequence Sequence) {
	CheckDB(db)

	_, err := db.Exec(`
		INSERT INTO weights(sequence, weight) 
		VALUES(?, ?)
		ON CONFLICT (sequence) DO UPDATE SET weight = excluded.weight
		`, sequence.Text, sequence.Weight)
	checkErr(err)
	log.Println("Added sequence:", sequence.Text)
}
