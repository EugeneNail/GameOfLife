package internal

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

func (game *Game) drawCellBorders(image *ebiten.Image, row int, column int) {
	if !game.isPaused {
		return
	}

	vector.DrawFilledRect(image,
		float32(column*game.cellSize),
		float32(row*game.cellSize),
		float32(game.cellSize),
		float32(game.cellSize),
		game.borderColor,
		false,
	)
}

func (game *Game) drawCell(image *ebiten.Image, row int, column int) {
	var cellColor color.Color
	if game.grid[row][column] {
		cellColor = color.White
	} else {
		cellColor = color.Black
	}

	vector.DrawFilledRect(image,
		float32(column*game.cellSize+1),
		float32(row*game.cellSize+1),
		float32(game.cellSize)-1,
		float32(game.cellSize)-1,
		cellColor,
		false,
	)
}

func (game *Game) drawGridBorders(image *ebiten.Image) {
	x1 := float32((game.columns) * game.cellSize)
	y1 := float32((game.rows) * game.cellSize)

	// top
	vector.StrokeLine(image, 0, 0, x1, 0, 1, game.borderColor, false)

	// right
	vector.StrokeLine(image, x1, 0, x1, y1, 1, game.borderColor, false)

	// bottom
	vector.StrokeLine(image, 0, y1, x1, y1, 1, game.borderColor, false)

	// left
	vector.StrokeLine(image, 0, 0, 0, y1, 1, game.borderColor, false)
}

func (game *Game) drawInfo(image *ebiten.Image) {
	height := float32(210)
	width := float32(150)
	xOffset := float32(10)

	_, screenHeight := ebiten.Monitor().Size()
	yOffset := (float32(screenHeight) - height) / 2

	vector.DrawFilledRect(image, xOffset, yOffset, width, height, color.NRGBA{R: 255, G: 255, B: 255, A: 30}, false)

	info := fmt.Sprintf(`
Game of life

Generation: %d
Population: %d
Speed: %d

LMB: paint a cell
RMB: pause
MMB: reset
Escape: exit
`,
		game.generation,
		game.population,
		game.Speed,
	)

	text.Draw(image, info, game.font, int(20+xOffset), int(20+yOffset), color.White)
}
