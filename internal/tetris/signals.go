package tetris

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// onKeyPressed : The onKeyPressed signal handler
func (t *Tetris) onKeyPressed(_ *gtk.ApplicationWindow, e *gdk.Event) {
	key := gdk.EventKeyNewFromEvent(e)

	switch key.KeyVal() {
	case 97: // Button "A" => Move tetromino left
		if !t.checkSideBlock(true) {
			t.falling.x -= 1
		}
	case 113: // Button "Q" => Quit game
		t.quitGame()
	case 115: // Button "S" => Rotate tetromino
		// Rotate every element except tetromino number 4 (the square)
		if t.falling.tetro.id != 4 {
			t.rotateTetromin(&t.falling.tetro)
		}
	case 100: // Button "D" => Move tetromino right
		if !t.checkSideBlock(false) {
			t.falling.x += 1
		}
	case 120: // Button "X" => Move tetromino down
		// TODO : Speed up tetromino
	}
	t.da.QueueDraw()
}

// onDraw : The onDraw signal handler
func (t *Tetris) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	t.drawBackground(da, ctx)
	t.drawPlayground(da, ctx)
	t.drawFallenTetrominos(da, ctx)
	t.drawFallingTetromino(da, ctx, t.falling.tetro)
}

//
// HELPER FUNCTIONS
//

// drawBackground : Draws the background
func (t *Tetris) drawBackground(da *gtk.DrawingArea, ctx *cairo.Context) {
	width := float64(da.GetAllocatedWidth())
	height := float64(da.GetAllocatedHeight())
	ctx.SetSourceRGB(0.4, 0.4, 1)
	ctx.Rectangle(0, 0, width, height)
	ctx.Fill()
}

// drawPlayground : Draws the playground background
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

// drawFallenTetrominos : Draws the tetrominos the have already fallen to the ground
func (t *Tetris) drawFallenTetrominos(da *gtk.DrawingArea, ctx *cairo.Context) {
	for y := 0; y < 20; y++ {
		for x := 0; x < 10; x++ {
			idx := t.playground[y][x]
			if idx > 0 {
				left, top := coordsToScreenCoords(x, y)
				t.drawBlock(da, ctx, tetrominos[idx-1].color, left, top)
			}
		}
	}
}

// drawFallingTetromino : Draws the currently falling tetromino
func (t *Tetris) drawFallingTetromino(da *gtk.DrawingArea, ctx *cairo.Context, tetro tetromino) {
	left, top := coordsToScreenCoords(t.falling.x, t.falling.y)

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !tetro.blocks[y][x] {
				continue
			}
			if t.falling.y-y > 19 {
				return
			}
			t.drawBlock(da, ctx, tetro.color, left+float64(x)*blockWidth, top+float64(y)*blockHeight)
		}
	}
}

// drawBlock : Draws a single block
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
