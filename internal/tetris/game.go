package tetris

import (
	"time"

	"github.com/hultan/tetris/internal/randomizer"
)

type game struct {
	speed    time.Duration
	isActive bool
	field    [fieldHeight][fieldWidth]int
	falling  fallingTetromino
	ticker   ticker
	rand     *randomizer.Randomizer
}

type fallingTetromino struct {
	tetro tetromino
	x, y  int
}

type ticker struct {
	tickerQuit chan struct{}
	ticker     *time.Ticker
}

// Drop a new Tetromino
func (g *game) nextTetromino() {
	g.falling.tetro = tetrominos[g.rand.Next()]
	g.falling.y = fieldVisibleHeight + tetrominoHeight - 1
	g.falling.x = (fieldWidth - tetrominoWidth) / 2
	g.speed -= 10

	// TODO : Remove, debug only
	g.rand.Print()
}

// Rotate the 4x4 tetromino array 90 degrees
// https://www.geeksforgeeks.org/rotate-a-matrix-by-90-degree-in-clockwise-direction-without-using-any-extra-space/
func (g *game) rotateTetromino(tetro *tetromino) {
	// Don't bother rotating the square
	if g.falling.tetro.id == 4 {
		return
	}

	cx, cy := g.falling.tetro.getRotationCenter()

	g.rotateTetrominoWithoutCheck(tetro)

	if g.checkOverlapping() {
		// Replace with counterclockwise rotation
		g.rotateTetrominoWithoutCheck(tetro)
		g.rotateTetrominoWithoutCheck(tetro)
		g.rotateTetrominoWithoutCheck(tetro)
		return
	}

	xx, yy := g.falling.tetro.getRotationCenter()

	g.falling.x += cx - xx
	g.falling.y -= cy - yy
}

func (g *game) rotateTetrominoWithoutCheck(tetro *tetromino) {
	// Rotate the tetromino array 90 degrees
	for y := 0; y < tetrominoHeight/2; y++ {
		for x := y; x < tetrominoWidth-y-1; x++ {
			tmp := tetro.blocks[y][x]
			tetro.blocks[y][x] = tetro.blocks[tetrominoHeight-1-x][y]
			tetro.blocks[tetrominoHeight-1-x][y] = tetro.blocks[tetrominoHeight-1-y][tetrominoWidth-1-x]
			tetro.blocks[tetrominoHeight-1-y][tetrominoWidth-1-x] = tetro.blocks[x][tetrominoWidth-1-y]
			tetro.blocks[x][tetrominoWidth-1-y] = tmp
		}
	}
}

func (g *game) checkOverlapping() bool {
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			// check if it is a block or not
			if g.falling.tetro.blocks[y][x] == 0 {
				continue
			}

			xx := g.falling.x + x
			yy := g.falling.y - y
			if xx < 0 || xx >= fieldWidth || yy < 0 {
				return true
			}
			if g.field[yy][xx] > 0 {
				return true
			}
		}
	}

	return false
}

func (g *game) checkFieldLimits(x, y int) bool {
	if x < 0 || x >= fieldWidth || y < 0 {
		return false
	}
	return g.field[y][x] > 0
}

func (g *game) checkFieldSides(left bool) bool {
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			// check if it is a block or not
			if g.falling.tetro.blocks[y][x] == 0 {
				continue
			}

			// check if left wall is blocking or
			// if a piece is blocking to the left
			if left && (g.falling.x+x == 0 || g.checkFieldLimits(g.falling.x+x-1, g.falling.y-y)) {
				return true
			}

			// check if right wall is blocking or
			// if a piece is blocking to the right
			if !left && (g.falling.x+x == fieldWidth-1 || g.checkFieldLimits(g.falling.x+x+1, g.falling.y-y)) {
				return true
			}
		}
	}

	return false
}

func (g *game) checkFieldBottom() bool {
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			// check if it is a block or not
			if g.falling.tetro.blocks[y][x] == 0 {
				continue
			}

			// check if floor is blocking
			if g.falling.y-y == 0 || g.checkFieldLimits(g.falling.x+x, g.falling.y-y-1) {
				g.moveFallingToField()
				return true
			}
		}
	}

	return false
}

func (g *game) moveFallingToField() {
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			if g.falling.tetro.blocks[y][x] > 0 {
				g.field[g.falling.y-y][g.falling.x+x] = g.falling.tetro.id
			}
		}
	}
	g.removeCompletedRows()
}

func (g *game) removeCompletedRows() {
	for y := 0; y < fieldHeight; y++ {
		rowComplete := true
		for x := 0; x < fieldWidth; x++ {
			if g.field[y][x] == 0 {
				rowComplete = false
				continue
			}
		}
		if rowComplete {
			g.deleteFieldRow(y)
			y -= 1
		}
	}
}

func (g *game) deleteFieldRow(d int) {
	for y := d; y < fieldHeight; y++ {
		for x := 0; x < fieldWidth; x++ {
			// To delete a row, copy the value from the row above
			// except for the top row, who should have zeroes
			if y == fieldHeight-1 {
				g.field[y][x] = 0
			} else {
				g.field[y][x] = g.field[y+1][x]
			}
		}
	}
}

func (g *game) dropTetrominoToField() {
	for !g.checkFieldBottom() {
		g.falling.y -= 1
	}
}
