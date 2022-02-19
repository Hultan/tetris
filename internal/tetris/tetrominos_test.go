package tetris

import (
	"fmt"
	"testing"
)

func Test_tetromino_getRotationCenter(t1 *testing.T) {
	tetro := tetrominos[2]
	g := &game{}

	dumpTetro(tetro)
	g.rotateTetromino(&tetro)
	dumpTetro(tetro)
	g.rotateTetromino(&tetro)
	dumpTetro(tetro)
	g.rotateTetromino(&tetro)
	dumpTetro(tetro)
}

func dumpTetro(tetro tetromino) {
	for i := 0; i < tetrominoHeight; i++ {
		s := ""
		for j := 0; j < tetrominoWidth; j++ {
			s += fmt.Sprintf("%d", tetro.blocks[i][j])
		}
		fmt.Println(s)
	}
	fmt.Println()
}
