package tetris

func (t *Tetris) checkBlockBottomSide() bool {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			// check if it is a block or not
			if !falling.blocks[y][x] {
				continue
			}

			// check if floor is blocking
			if posY-y == 0 || t.checkPlayground(posX+x, posY-y-1) {
				t.moveFallingToFallen()
				return true
			}
		}
	}

	return false
}

func (t *Tetris) checkSideBlock(left bool) bool {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			// check if it is a block or not
			if !falling.blocks[y][x] {
				continue
			}

			// check if left wall is blocking or
			// if a piece is blocking to the left
			if left && (posX+x == 0 || t.checkPlayground(posX+x-1, posY-y)) {
				return true
			}

			// check if right wall is blocking or
			// if a piece is blocking to the right
			if !left && (posX+x == 9 || t.checkPlayground(posX+x+1, posY-y)) {
				return true
			}
		}
	}

	return false
}

func (t *Tetris) checkPlayground(x, y int) bool {
	if x < 0 || x > 9 || y < 0 {
		return false
	}
	return playground[y][x] > 0
}

func (t *Tetris) moveFallingToFallen() {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if falling.blocks[y][x] {
				playground[posY-y][posX+x] = falling.id
			}
		}
	}
	t.removeCompleteRows()
}

func (t *Tetris) removeCompleteRows() {
	for y := 0; y < 25; y++ {
		rowComplete := true
		for x := 0; x < 9; x++ {
			if playground[y][x] == 0 {
				rowComplete = false
			}
		}
		if rowComplete {
			t.deleteRow(y)
		}
	}
}

func (t *Tetris) deleteRow(d int) {
	for y := d; y < 25; y++ {
		for x := 0; x < 9; x++ {
			if y == 24 {
				playground[y][x] = 0
			} else {
				playground[y][x] = playground[y+1][x]
			}
		}
	}
}
