package tetris

func (t *Tetris) checkBlockBottomSide() bool {
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			// check if it is a block or not
			if !t.falling.tetro.blocks[y][x] {
				continue
			}

			// check if floor is blocking
			if t.falling.y-y == 0 || t.checkPlayfield(t.falling.x+x, t.falling.y-y-1) {
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
			if left && (t.falling.x+x == 0 || t.checkPlayfield(t.falling.x+x-1, t.falling.y-y)) {
				return true
			}

			// check if right wall is blocking or
			// if a piece is blocking to the right
			if !left && (t.falling.x+x == playfieldWidth-1 || t.checkPlayfield(t.falling.x+x+1, t.falling.y-y)) {
				return true
			}
		}
	}

	return false
}

func (t *Tetris) checkPlayfield(x, y int) bool {
	if x < 0 || x >= playfieldWidth || y < 0 {
		return false
	}
	return t.playfield[y][x] > 0
}

func (t *Tetris) moveFallingToFallen() {
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			if t.falling.tetro.blocks[y][x] {
				t.playfield[t.falling.y-y][t.falling.x+x] = t.falling.tetro.id
			}
		}
	}
	t.removeCompleteRows()
}

func (t *Tetris) removeCompleteRows() {
	for y := 0; y < playfieldHeight; y++ {
		rowComplete := true
		for x := 0; x < playfieldWidth; x++ {
			if t.playfield[y][x] == 0 {
				rowComplete = false
				continue
			}
		}
		if rowComplete {
			t.deleteRow(y)
			y -= 1
		}
	}
}

func (t *Tetris) deleteRow(d int) {
	for y := d; y < playfieldHeight; y++ {
		for x := 0; x < playfieldWidth; x++ {
			// To delete a row, copy the value from the row above
			// except for the top row, who should have zeroes
			if y == playfieldHeight-1 {
				t.playfield[y][x] = 0
			} else {
				t.playfield[y][x] = t.playfield[y+1][x]
			}
		}
	}
}
