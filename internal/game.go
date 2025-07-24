package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"math"
	"os"
)

type Game struct {
	isPaused                   bool
	frameCounter               int
	Speed                      int
	frameInterval              int
	grid                       [][]bool
	rows                       int
	columns                    int
	cellSize                   int
	leftMouseButtonWasPressed  bool
	rightMouseButtonWasPressed bool
	borderColor                color.RGBA
	font                       font.Face
	generation                 int
	population                 int
}

func NewGame(columns int, screenWidth int, screenHeight int) *Game {
	game := &Game{}
	game.importFont()

	game.grid = make([][]bool, 0)
	if columns == 0 {
		game.cellSize = 20
	} else {
		game.cellSize = int(math.Floor(float64(screenWidth / columns)))
	}
	game.columns = int(math.Floor(float64(screenWidth / game.cellSize)))
	game.rows = int(math.Floor(float64(screenHeight / game.cellSize)))
	game.grid = createGrid(game.rows, game.columns)
	game.Speed = 10
	game.frameInterval = 2
	game.borderColor = color.RGBA{R: 33, G: 33, B: 33, A: 255}
	game.isPaused = true

	return game
}

func (game *Game) importFont() {
	fontBytes, err := os.ReadFile("./assets/font/Roboto-Medium.ttf")
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}

	game.font, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 14,
		DPI:  72,
	})

	if err != nil {
		panic(err)
	}
}

func (game *Game) Layout(_ int, _ int) (int, int) {
	return game.cellSize * game.columns, game.cellSize * game.rows
}

func (game *Game) Update() error {
	game.frameCounter++
	game.handleCellPainting()
	game.handlePausing()
	game.handleExiting()
	game.handleSpeeding()
	game.handleResetting()

	if game.isPaused || game.frameCounter%game.frameInterval != 0 {
		return nil
	}

	newGrid := createGrid(game.rows, game.columns)
	game.population = 0

	for row, cells := range newGrid {
		for column, _ := range cells {
			aliveNeighbors := game.countAliveNeighbors(row, column)
			newGrid[row][column] = game.willCellLive(row, column, aliveNeighbors)
			game.population += boolToInt(newGrid[row][column])
		}
	}

	game.grid = newGrid
	game.generation++

	return nil
}

func (game *Game) countAliveNeighbors(row int, column int) int {
	var aliveNeighbors int

	// top left
	aliveNeighbors += boolToInt(game.isCellAlive(row, column, -1, -1))

	// top
	aliveNeighbors += boolToInt(game.isCellAlive(row, column, -1, 0))

	// top right
	aliveNeighbors += boolToInt(game.isCellAlive(row, column, -1, +1))

	// right
	aliveNeighbors += boolToInt(game.isCellAlive(row, column, 0, +1))

	// bottom right
	aliveNeighbors += boolToInt(game.isCellAlive(row, column, +1, +1))

	// bottom
	aliveNeighbors += boolToInt(game.isCellAlive(row, column, +1, 0))

	// bottom left
	aliveNeighbors += boolToInt(game.isCellAlive(row, column, +1, -1))

	// left
	aliveNeighbors += boolToInt(game.isCellAlive(row, column, 0, -1))

	return aliveNeighbors
}

func (game *Game) willCellLive(row int, column int, aliveNeighbors int) bool {
	isAlive := game.grid[row][column]

	if isAlive && aliveNeighbors <= 1 {
		return false
	}

	if isAlive && aliveNeighbors >= 4 {
		return false
	}

	if isAlive && aliveNeighbors >= 2 && aliveNeighbors <= 3 {
		return true
	}

	if !isAlive && aliveNeighbors == 3 {
		return true
	}

	return false
}

func (game *Game) isCellAlive(row int, column int, rowOffset int, columnOffset int) bool {
	rowEdge := game.rows - 1
	calculatedRow := row + rowOffset
	if calculatedRow > rowEdge {
		calculatedRow = 0
	} else if calculatedRow < 0 {
		calculatedRow = rowEdge
	}

	columnEdge := game.columns - 1
	calculatedColumn := column + columnOffset
	if calculatedColumn > columnEdge {
		calculatedColumn = 0
	} else if calculatedColumn < 0 {
		calculatedColumn = columnEdge
	}

	return game.grid[calculatedRow][calculatedColumn]
}

func (game *Game) Draw(image *ebiten.Image) {
	for row, cells := range game.grid {
		for column, _ := range cells {
			game.drawCellBorders(image, row, column)
			game.drawCell(image, row, column)
		}
	}

	game.drawGridBorders(image)
	game.drawInfo(image)
}
