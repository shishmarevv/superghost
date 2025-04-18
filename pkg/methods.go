package superghost

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func (s *Sequence) Update() {
	db := Open()
	CheckDB(db)
	defer Shut(db)

	rows, err := db.Query(`
		SELECT weight
		FROM weights
		WHERE sequence = ?
		`, s.Text)
	checkErr(err)
	defer rows.Close()

	if rows.Next() {
		for rows.Next() {
			err = rows.Scan(&s.Weight)
			checkErr(err)
		}
	} else {
		s.Weight = 1.0
		AddSequence(db, *s)
	}

}

func (s *Sequence) Add(symbol string, direction bool) {
	if len(symbol) != 1 {
		log.Fatal("Wrong symbol length")
	}
	if direction {
		//left
		s.Text = symbol + s.Text
	} else {
		//right
		s.Text = s.Text + symbol
	}
	s.Update()
}

func (g *Game) Update(sequence Sequence) {
	g.Turn = !g.Turn
	g.Word = sequence.Text
	db := Open()
	if !IsInWord(db, g.Word) || IsWord(db, g.Word) {
		g.On = false
	}
}

func (g *Game) Start() {
	g.Turn = true
	g.On = true
	g.Word = ""
}
