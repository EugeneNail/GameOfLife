package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"os"
)

func (game *Game) handleCellPainting() {
	if !game.isPaused {
		return
	}

	isMousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if isMousePressed && !game.leftMouseButtonWasPressed {
		x, y := ebiten.CursorPosition()
		row := y / game.cellSize
		column := x / game.cellSize

		if column < game.columns && row < game.rows {
			game.grid[row][column] = !game.grid[row][column]
		}
	}

	game.leftMouseButtonWasPressed = isMousePressed
}

func (game *Game) handlePausing() {
	rightMouseWasPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	if rightMouseWasPressed && !game.rightMouseButtonWasPressed {
		game.isPaused = !game.isPaused
	}

	game.rightMouseButtonWasPressed = rightMouseWasPressed
}

func (game *Game) handleSpeeding() {
	_, yOffset := ebiten.Wheel()

	if yOffset > 0 && game.Speed+game.frameInterval <= game.frameInterval*25 {
		game.Speed += game.frameInterval
	} else if yOffset < 0 && game.Speed-game.frameInterval >= 10 {
		game.Speed -= game.frameInterval
	}

	ebiten.SetTPS(game.Speed)
}

func (game *Game) handleExiting() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(1)
	}
}

func (game *Game) handleResetting() {
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		return
	}

	game.isPaused = true
	game.population = 0
	game.generation = 0
	game.Speed = 10
	game.grid = createGrid(game.rows, game.columns)
}
