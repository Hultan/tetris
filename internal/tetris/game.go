package tetris

import (
	"time"

	"github.com/hultan/tetris/internal/randomizer"
)

type game struct {
	speed     time.Duration
	isActive  bool
	playfield [playfieldHeight][playfieldWidth]int
	falling   fallingTetromino
	ticker    ticker
	rand      *randomizer.Randomizer
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
func (g *game) createNewFallingTetromino() {
	g.falling.tetro = tetrominos[g.rand.Next()]
	g.falling.y = playfieldVisibleHeight + tetrominoHeight - 1
	g.falling.x = (playfieldWidth - tetrominoWidth) / 2
	g.speed -= 10

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
			if xx < 0 || xx >= playfieldWidth || yy < 0 {
				return true
			}
			if g.playfield[yy][xx] > 0 {
				return true
			}
		}
	}

	return false
}

func (g *game) checkPlayfieldLimits(x, y int) bool {
	if x < 0 || x >= playfieldWidth || y < 0 {
		return false
	}
	return g.playfield[y][x] > 0
}

func (g *game) checkPlayfieldSides(left bool) bool {
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			// check if it is a block or not
			if g.falling.tetro.blocks[y][x] == 0 {
				continue
			}

			// check if left wall is blocking or
			// if a piece is blocking to the left
			if left && (g.falling.x+x == 0 || g.checkPlayfieldLimits(g.falling.x+x-1, g.falling.y-y)) {
				return true
			}

			// check if right wall is blocking or
			// if a piece is blocking to the right
			if !left && (g.falling.x+x == playfieldWidth-1 || g.checkPlayfieldLimits(g.falling.x+x+1, g.falling.y-y)) {
				return true
			}
		}
	}

	return false
}

func (g *game) checkPlayfieldBottom() bool {
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			// check if it is a block or not
			if g.falling.tetro.blocks[y][x] == 0 {
				continue
			}

			// check if floor is blocking
			if g.falling.y-y == 0 || g.checkPlayfieldLimits(g.falling.x+x, g.falling.y-y-1) {
				g.moveFallingToPlayfield()
				return true
			}
		}
	}

	return false
}

func (g *game) moveFallingToPlayfield() {
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			if g.falling.tetro.blocks[y][x] > 0 {
				g.playfield[g.falling.y-y][g.falling.x+x] = g.falling.tetro.id
			}
		}
	}
	g.removeCompletePlayfieldRows()
}

func (g *game) removeCompletePlayfieldRows() {
	for y := 0; y < playfieldHeight; y++ {
		rowComplete := true
		for x := 0; x < playfieldWidth; x++ {
			if g.playfield[y][x] == 0 {
				rowComplete = false
				continue
			}
		}
		if rowComplete {
			g.deletePlayfieldRow(y)
			y -= 1
		}
	}
}

func (g *game) deletePlayfieldRow(d int) {
	for y := d; y < playfieldHeight; y++ {
		for x := 0; x < playfieldWidth; x++ {
			// To delete a row, copy the value from the row above
			// except for the top row, who should have zeroes
			if y == playfieldHeight-1 {
				g.playfield[y][x] = 0
			} else {
				g.playfield[y][x] = g.playfield[y+1][x]
			}
		}
	}
}

func (g *game) dropTetrominoToPlayfield() {
	for !g.checkPlayfieldBottom() {
		g.falling.y -= 1
	}
}
