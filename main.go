package main

import (
	"log"

	"bataille/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(g.Width, g.Height)
	ebiten.SetWindowTitle("Bataille")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

	return
}
