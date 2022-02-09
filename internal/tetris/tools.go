package tetris

import (
	"fmt"
	"math/rand"
)

func col(c uint32) float64 {
	return float64(c) / 65535
}

func rotate(t *tetromino) {
	// // Traverse each cycle
	// for (int i = 0; i < N / 2; i++) {
	// 	for (int j = i; j < N - i - 1; j++) {
	//
	// 		// Swap elements of each cycle
	// 		// in clockwise direction
	// 		int temp = a[i][j];
	// 		a[i][j] = a[N - 1 - j][i];
	// 		a[N - 1 - j][i] = a[N - 1 - i][N - 1 - j];
	// 		a[N - 1 - i][N - 1 - j] = a[j][N - 1 - i];
	// 		a[j][N - 1 - i] = temp;
	// 	}
	// }
	for i := 0; i < 5/2; i++ {
		for j := 0; j < 5-i-1; j++ {
			tmp := t.blocks[i][j]
			t.blocks[i][j] = t.blocks[5-1-j][i]
			t.blocks[5-1-j][i] = t.blocks[5-1-i][5-1-j]
			t.blocks[5-1-i][5-1-j] = t.blocks[j][5-1-i]
			t.blocks[j][5-1-i] = tmp
		}
	}
	// for i := 0; i < 4; i++ {
	// 	for j := i + 1; j < 5; j++ {
	// 		t.blocks[i][j], t.blocks[j][i] = t.blocks[j][i], t.blocks[i][j]
	// 	}
	// }
}

func newFallingTetromino() {
	r := rand.Intn(7)
	falling = tetrominos[r]
	posY = 24
	posX = 3
}

func coordsToScreenCoords(x, y int) (float64, float64) {
	return float64(leftBorder + x*blockWidth), float64(topBorder + (19-y)*blockHeight)
}

func printPlayground() {
	fmt.Println()
	fmt.Println("----------------")
	fmt.Println()
	for r := 0; r < 25; r++ {
		fmt.Println(playground[r])
	}
}
