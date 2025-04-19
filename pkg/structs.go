package superghost

import "bufio"

type Sequence struct {
	Text   string
	Weight float64
}

type Game struct {
	Word   string
	On     bool
	Turn   bool //First player = true; Second player = false
	Source Input
}

type Input interface {
	GetSymbol() (string, error)
	GetDirection() (bool, error)
}

type CMDinput struct {
	reader *bufio.Reader
}

type Output interface {
	Out(string)
	Winner(*Game)
}

type CMDoutput struct {
	writer *bufio.Writer
}
