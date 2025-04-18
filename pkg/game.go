package superghost

import (
	"fmt"
	"log"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Player_versus_player(InputProvider Input) {
	log.Println("Player versus player")

	var game = NewGame(InputProvider)
	game.Start()

	var buffer Sequence
	buffer.Text = ""
	buffer.Weight = 1.0

	for game.On {

		symbol := game.Source.GetSymbol()

		if symbol == "" {
			continue
		}
	}
}
