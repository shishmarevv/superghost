package superghost

import (
	"bufio"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
	"unicode"
)

func (s *Sequence) Update() {
	db := Open()
	CheckDB(db)
	defer Shut(db)

	rows, err := db.Query(`
		SELECT weight
		FROM weights
		WHERE sequence LIKE ?
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

func NewCMDoutput() *CMDoutput {
	return &CMDoutput{
		writer: bufio.NewWriter(os.Stdout),
	}
}

func (c *CMDinput) GetSymbol() (string, error) {
	fmt.Println("Letter:")
	input, err := c.reader.ReadString('\n')
	CheckErr(err)

	input = strings.TrimSpace(input)
	if len(input) != 1 {
		return "", errors.New("wrong symbol length")
	} else if !unicode.IsLetter(rune(input[0])) {
		return "", errors.New("invalid symbol")
	}
	return strings.ToLower(input), nil
}

func (c *CMDinput) GetDirection() (bool, error) {
	fmt.Println("Direction:")
	input, err := c.reader.ReadString('\n')
	CheckErr(err)

	input = strings.TrimSpace(input)
	if input == "left" {
		return true, nil
	} else if input == "right" {
		return false, nil
	}
	return false, errors.New("not a direction")
}

func (c *CMDoutput) Out(message string) {
	_, err := c.writer.WriteString("Current sequence:" + message)
	CheckErr(err)
	err = c.writer.Flush()
	CheckErr(err)
}

func (c *CMDoutput) Winner(game *Game) {
	if !game.Turn {
		_, err := c.writer.WriteString("First player won")
		CheckErr(err)
	} else {
		_, err := c.writer.WriteString("Second player won")
		CheckErr(err)
	}
}
