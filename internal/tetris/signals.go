package tetris

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func (t *Tetris) onKeyPressed(_ *gtk.ApplicationWindow, e *gdk.Event) {
	key := gdk.EventKeyNewFromEvent(e)

	switch key.KeyVal() {
	case 97: // Button "A"
		if !t.checkBlockLeftSide() {
			posX -= 1
		}
	case 113: // Button "Q"
		if isPlaying {
			close(quitChannel) // Stop ticker
		}
		t.w.Close() // Close window
	case 115: // Button "S"
		rotate(&falling)
	case 100: // Button "D"
		if !t.checkBlockRightSide() {
			posX += 1
		}
	case 120: // Button "X"
		// TODO : Speed up tetromino
	}
	t.da.QueueDraw()
}

func (t *Tetris) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	t.drawBackground(da, ctx)
	t.drawPlayground(da, ctx)
	t.drawFallenTetrominos(da, ctx)
	t.drawFallingTetromino(da, ctx, falling)
}

func (t *Tetris) drawBackground(da *gtk.DrawingArea, ctx *cairo.Context) {
	width := float64(da.GetAllocatedWidth())
	height := float64(da.GetAllocatedHeight())
	ctx.SetSourceRGB(0.4, 0.4, 1)
	ctx.Rectangle(0, 0, width, height)
	ctx.Fill()
}

func (t *Tetris) drawFallingTetromino(da *gtk.DrawingArea, ctx *cairo.Context, tetro tetromino) {
	left, top := coordsToScreenCoords(posX, posY)

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !tetro.blocks[y][x] {
				continue
			}
			t.drawBlock(da, ctx, tetro.color, left+float64(x)*blockWidth, top+float64(y)*blockHeight)
		}
	}
}

func (t *Tetris) drawBlock(_ *gtk.DrawingArea, ctx *cairo.Context, c color.Color, left, top float64) {
	// Fill block in the correct color
	r, g, b, a := c.RGBA()
	ctx.SetSourceRGBA(col(r), col(g), col(b), col(a))
	ctx.Rectangle(left, top, blockWidth, blockHeight)
	ctx.Fill()

	// Draw black border around block
	ctx.SetSourceRGBA(0, 0, 0, 1)
	ctx.SetLineWidth(1)
	ctx.Rectangle(left, top, blockWidth, blockHeight)
	ctx.Stroke()
}

func (t *Tetris) drawFallenTetrominos(da *gtk.DrawingArea, ctx *cairo.Context) {
	for y := 0; y < 20; y++ {
		for x := 0; x < 10; x++ {
			idx := playground[y][x]
			if idx > 0 {
				left, top := coordsToScreenCoords(x, y)
				t.drawBlock(da, ctx, tetrominos[idx-1].color, left, top)
			}
		}
	}
}

func (t *Tetris) drawPlayground(da *gtk.DrawingArea, ctx *cairo.Context) {
	ctx.SetSourceRGBA(1, 1, 1, 1)
	ctx.Rectangle(leftBorder, topBorder, 10*blockWidth, 20*blockHeight)
	ctx.Fill()

	ctx.SetSourceRGBA(0.5, 0.5, 0.5, 1)
	ctx.SetLineWidth(1)
	for i := 0; i < 10; i++ {
		// Vertical lines
		ctx.MoveTo(float64(leftBorder+(i+1)*blockWidth), topBorder)
		ctx.LineTo(float64(leftBorder+(i+1)*blockWidth), topBorder+20*blockHeight)
		ctx.Stroke()
	}
	for i := 0; i < 20; i++ {
		// Horizontal lines
		ctx.MoveTo(leftBorder, float64(topBorder+(i+1)*blockHeight))
		ctx.LineTo(leftBorder+10*blockWidth, float64(topBorder+(i+1)*blockHeight))
		ctx.Stroke()
	}
}

func (t *Tetris) checkBlock() bool {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			// check if floor is blocking
			if falling.blocks[y][x] && posY-y == 0 {
				t.moveFallingToFallen()
				return true
			}

			// check if fallen tetros is blocking
			if falling.blocks[y][x] && playground[posY-y-1][posX+x] > 0 {
				t.moveFallingToFallen()
				return true
			}
		}
	}

	return false
}

func (t *Tetris) checkBlockLeftSide() bool {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			// check if left wall is blocking
			if falling.blocks[y][x] && posX+x == 0 {
				return true
			}

			// check if a piece is blocking to the left
			if falling.blocks[y][x] && playground[posY-y][posX+x-1] > 0 {
				return true
			}
		}
	}

	return false
}

func (t *Tetris) checkBlockRightSide() bool {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			// check if right wall is blocking
			if falling.blocks[y][x] && posX+x == 9 {
				return true
			}

			// check if a piece is blocking to the right
			if falling.blocks[y][x] && playground[posY-y][posX+x+1] > 0 {
				return true
			}
		}
	}

	return false
}

func (t *Tetris) moveFallingToFallen() {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if falling.blocks[y][x] {
				playground[posY-y][posX+x] = falling.id
			}
		}
	}
	printPlayground()
}
