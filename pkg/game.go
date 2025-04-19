package superghost

import (
	"log"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Player_versus_player(InputProvider Input, OutputProvider Output) {
	log.Println("Player versus player")

	var game = NewGame(InputProvider)
	game.Start()

	var buffer Sequence
	buffer.Text = ""
	buffer.Weight = 1.0

	for game.On {
		if game.Turn {
			log.Println("First player")
		} else {
			log.Println("Second player")
		}
		direction, err := game.Source.GetDirection()
		if err != nil {
			log.Println(err)
			continue
		}
		symbol, err := game.Source.GetSymbol()
		if err != nil {
			log.Println(err)
			continue
		}
		buffer.Add(symbol, direction)
		buffer.Update()
		game.Update(buffer)
		OutputProvider.Out(buffer.Text)
	}
	OutputProvider.Winner(game)
}
