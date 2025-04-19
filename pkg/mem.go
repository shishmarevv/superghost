package superghost

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Open() *sql.DB {
	root := FindProjectRoot()
	dbPath := filepath.Join(root, "memory", "mem.db")
	db, err := sql.Open("sqlite3", dbPath)
	CheckErr(err)

	log.Println("Database opened")

	return db
}

func Shut(db *sql.DB) {
	err := db.Close()
	CheckErr(err)

	log.Println("Database closed")
}

func CheckDB(db *sql.DB) bool {
	err := db.Ping()
	CheckErr(err)

	return true
}

func AddWord(db *sql.DB, word string) {
	CheckDB(db)

	_, err := db.Exec(`
		INSERT INTO words(word) 
		VALUES(?)
		`, strings.ToLower(word))
	CheckErr(err)

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
	CheckErr(err)

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
		WHERE word LIKE ?
		`, word)
	CheckErr(err)

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
		`, strings.ToLower(sequence.Text), sequence.Weight)
	CheckErr(err)
	log.Println("Added sequence:", sequence.Text)
}

func FindProjectRoot() string {
	dir, _ := os.Getwd()
	for {
		if fi, err := os.Stat(filepath.Join(dir, ".git")); err == nil && fi.IsDir() {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			log.Fatalf(".git not found in any parent of %s", dir)
		}
		dir = parent
	}
}
