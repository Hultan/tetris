package tetris

func (t *Tetris) checkPlayfieldLimits(x, y int) bool {
	if x < 0 || x >= playfieldWidth || y < 0 {
		return false
	}
	return t.game.playfield[y][x] > 0
}

func (t *Tetris) checkPlayfieldSides(left bool) bool {
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			// check if it is a block or not
			if !t.game.falling.tetro.blocks[y][x] {
				continue
			}

			// check if left wall is blocking or
			// if a piece is blocking to the left
			if left && (t.game.falling.x+x == 0 || t.checkPlayfieldLimits(t.game.falling.x+x-1, t.game.falling.y-y)) {
				return true
			}

			// check if right wall is blocking or
			// if a piece is blocking to the right
			if !left && (t.game.falling.x+x == playfieldWidth-1 || t.checkPlayfieldLimits(t.game.falling.x+x+1, t.game.falling.y-y)) {
				return true
			}
		}
	}

	return false
}

func (t *Tetris) checkPlayfieldBottom() bool {
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			// check if it is a block or not
			if !t.game.falling.tetro.blocks[y][x] {
				continue
			}

			// check if floor is blocking
			if t.game.falling.y-y == 0 || t.checkPlayfieldLimits(t.game.falling.x+x, t.game.falling.y-y-1) {
				t.moveFallingToPlayfield()
				return true
			}
		}
	}

	return false
}

func (t *Tetris) moveFallingToPlayfield() {
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			if t.game.falling.tetro.blocks[y][x] {
				t.game.playfield[t.game.falling.y-y][t.game.falling.x+x] = t.game.falling.tetro.id
			}
		}
	}
	t.removeCompletePlayfieldRows()
}

func (t *Tetris) removeCompletePlayfieldRows() {
	for y := 0; y < playfieldHeight; y++ {
		rowComplete := true
		for x := 0; x < playfieldWidth; x++ {
			if t.game.playfield[y][x] == 0 {
				rowComplete = false
				continue
			}
		}
		if rowComplete {
			t.deletePlayfieldRow(y)
			y -= 1
		}
	}
}

func (t *Tetris) deletePlayfieldRow(d int) {
	for y := d; y < playfieldHeight; y++ {
		for x := 0; x < playfieldWidth; x++ {
			// To delete a row, copy the value from the row above
			// except for the top row, who should have zeroes
			if y == playfieldHeight-1 {
				t.game.playfield[y][x] = 0
			} else {
				t.game.playfield[y][x] = t.game.playfield[y+1][x]
			}
		}
	}
}

//
// // Debug function : Print the playfield
// func (t *Tetris) printPlayfield() {
// 	fmt.Println()
// 	fmt.Println("----------------")
// 	fmt.Println()
// 	for r := 0; r < playfieldHeight; r++ {
// 		fmt.Println(t.game.playfield[r])
// 	}
// }
