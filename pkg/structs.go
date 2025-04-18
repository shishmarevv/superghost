package superghost

type Sequence struct {
	Text   string
	Weight float64
}

type Game struct {
	Word string
	On   bool
	Turn bool //First player = true; Second player = false
}
