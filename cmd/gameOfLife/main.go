package main

import (
	"flag"
	"github.com/EugeneNail/GameOfLife/internal"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	columnsFlag := flag.Int("columns", 0, "defines number of columns per row")
	flag.Parse()
	width, height := ebiten.Monitor().Size()
	game := internal.NewGame(*columnsFlag, width, height)

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Game of life")
	ebiten.SetTPS(game.Speed)
	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
