package tetris

func (t *Tetris) checkBlockBottomSide() bool {
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			// check if it is a block or not
			if !t.falling.tetro.blocks[y][x] {
				continue
			}

			// check if floor is blocking
			if t.falling.y-y == 0 || t.checkPlayground(t.falling.x+x, t.falling.y-y-1) {
				t.moveFallingToFallen()
				return true
			}
		}
	}

	return false
}

func (t *Tetris) checkSideBlock(left bool) bool {
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			// check if it is a block or not
			if !t.falling.tetro.blocks[y][x] {
				continue
			}

			// check if left wall is blocking or
			// if a piece is blocking to the left
			if left && (t.falling.x+x == 0 || t.checkPlayground(t.falling.x+x-1, t.falling.y-y)) {
				return true
			}

			// check if right wall is blocking or
			// if a piece is blocking to the right
			if !left && (t.falling.x+x == playgroundWidth-1 || t.checkPlayground(t.falling.x+x+1, t.falling.y-y)) {
				return true
			}
		}
	}

	return false
}

func (t *Tetris) checkPlayground(x, y int) bool {
	if x < 0 || x >= playgroundWidth || y < 0 {
		return false
	}
	return t.playground[y][x] > 0
}

func (t *Tetris) moveFallingToFallen() {
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			if t.falling.tetro.blocks[y][x] {
				t.playground[t.falling.y-y][t.falling.x+x] = t.falling.tetro.id
			}
		}
	}
	t.removeCompleteRows()
}

func (t *Tetris) removeCompleteRows() {
	for y := 0; y < playgroundHeight; y++ {
		rowComplete := true
		for x := 0; x < playgroundWidth; x++ {
			if t.playground[y][x] == 0 {
				rowComplete = false
				continue
			}
		}
		if rowComplete {
			t.deleteRow(y)
		}
	}
}

func (t *Tetris) deleteRow(d int) {
	for y := d; y < playgroundHeight; y++ {
		for x := 0; x < playgroundWidth; x++ {
			// To delete a row, copy the value from the row above
			// except for the top row, who should have zeroes
			if y == playgroundHeight-1 {
				t.playground[y][x] = 0
			} else {
				t.playground[y][x] = t.playground[y+1][x]
			}
		}
	}
}
