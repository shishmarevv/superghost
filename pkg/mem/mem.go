package mem

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"superghost/pkg/structs"
)

func Open() *sql.DB {
	db, err := sql.Open("sqlite3", "./mem.db")
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

func AddSequence(db *sql.DB, sequence structs.Sequence) {
	CheckDB(db)

	_, err := db.Exec(`
		INSERT INTO weights(sequence, weight) 
		VALUES(?, ?)
		ON CONFLICT (sequence) DO UPDATE SET weight = excluded.weight
		`, sequence.Text, sequence.Weight)
	checkErr(err)
	log.Println("Added sequence:", sequence.Text)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
