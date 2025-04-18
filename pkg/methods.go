package superghost

import (
	"bufio"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
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
	CheckErr(err)
	defer rows.Close()

	if rows.Next() {
		for rows.Next() {
			err = rows.Scan(&s.Weight)
			CheckErr(err)
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

func NewGame(source Input) *Game {
	return &Game{
		Source: source,
	}
}

func NewCMDinput() *CMDinput {
	return &CMDinput{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (c *CMDinput) GetSymbol() string {
	fmt.Println("Letter:")
	input, err := c.reader.ReadString('\n')
	CheckErr(err)

	symbol := strings.TrimSpace(input)
	if len(symbol) == 0 {
		return ""
	}

	return symbol[:1]
}

func (c *CMDinput) GetDirection() bool {
	fmt.Println("Direction:")
	input, err := c.reader.ReadString('\n')
	CheckErr(err)

	input = strings.TrimSpace(input)
	if input == "left" {
		return true
	} else if input == "right" {
		return false
	}
	log.Fatal("Not a direction")
	return nil
}
