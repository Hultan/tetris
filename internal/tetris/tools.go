package tetris

import (
	"fmt"
	"math/rand"
)

func col(c uint32) float64 {
	return float64(c) / 65535
}

// Rotate the 5x5 tetromino array 90 degrees
func rotate(t *tetromino) {
	for i := 0; i < 5/2; i++ {
		for j := 0; j < 5-i-1; j++ {
			tmp := t.blocks[i][j]
			t.blocks[i][j] = t.blocks[5-1-j][i]
			t.blocks[5-1-j][i] = t.blocks[5-1-i][5-1-j]
			t.blocks[5-1-i][5-1-j] = t.blocks[j][5-1-i]
			t.blocks[j][5-1-i] = tmp
		}
	}
}

// Drop a new Tetromino
func newFallingTetromino() {
	r := rand.Intn(7)
	falling = tetrominos[r]
	posY = 24
	posX = 3
}

// Convert playground coords to screen coords
func coordsToScreenCoords(x, y int) (float64, float64) {
	return float64(leftBorder + x*blockWidth), float64(topBorder + (19-y)*blockHeight)
}

// Debug function : Print the playground
func printPlayground() {
	fmt.Println()
	fmt.Println("----------------")
	fmt.Println()
	for r := 0; r < 25; r++ {
		fmt.Println(playground[r])
	}
}

func adjustPosition() {
	min, max := 0, 9
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !falling.blocks[y][x] {
				continue
			}
			if posX+x < min {
				min = posX + x
			}
			if posX+x > max {
				max = posX + x
			}
		}
	}

	if min < 0 {
		posX += -min
	}
	if max > 9 {
		posX -= max - 9
	}
}
