package tetris

import (
	"math/rand"
)

func (t *Tetris) adjustPositionAfterRotate() {
	min, max := 0, playfieldWidth-1
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			if !t.game.falling.tetro.blocks[y][x] {
				continue
			}
			if t.game.falling.x+x < min {
				min = t.game.falling.x + x
			}
			if t.game.falling.x+x > max {
				max = t.game.falling.x + x
			}
		}
	}

	if min < 0 {
		t.game.falling.x += -min
	}
	if max > playfieldWidth-1 {
		t.game.falling.x -= max - (playfieldWidth - 1)
	}
}

// Drop a new Tetromino
func (t *Tetris) createNewFallingTetromino() {
	r := rand.Intn(tetrominoCount)
	t.game.falling.tetro = tetrominos[r]
	t.game.falling.y = playfieldHeight - 1
	t.game.falling.x = (playfieldWidth - tetrominoWidth) / 2
	t.game.speed -= 10
}

// Rotate the 4x4 tetromino array 90 degrees
// https://www.geeksforgeeks.org/rotate-a-matrix-by-90-degree-in-clockwise-direction-without-using-any-extra-space/
func (t *Tetris) rotateTetromino(tetro *tetromino) {
	for y := 0; y < tetrominoHeight/2; y++ {
		for x := y; x < tetrominoWidth-y-1; x++ {
			tmp := tetro.blocks[y][x]
			tetro.blocks[y][x] = tetro.blocks[tetrominoHeight-1-x][y]
			tetro.blocks[tetrominoHeight-1-x][y] = tetro.blocks[tetrominoHeight-1-y][tetrominoWidth-1-x]
			tetro.blocks[tetrominoHeight-1-y][tetrominoWidth-1-x] = tetro.blocks[x][tetrominoWidth-1-y]
			tetro.blocks[x][tetrominoWidth-1-y] = tmp
		}
	}

	t.adjustPositionAfterRotate()
}
