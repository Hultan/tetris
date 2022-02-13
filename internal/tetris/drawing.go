package tetris

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

// onDraw : The onDraw signal handler
func (t *Tetris) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	t.drawBackground(da, ctx)
	t.drawPlayfield(da, ctx)
	t.drawFallenTetrominos(da, ctx)
	t.drawFallingTetromino(da, ctx, t.game.falling.tetro)
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

// drawPlayfield : Draws the playfield background
func (t *Tetris) drawPlayfield(da *gtk.DrawingArea, ctx *cairo.Context) {
	ctx.SetSourceRGBA(1, 1, 1, 1)
	ctx.Rectangle(leftBorder, topBorder, 10*blockWidth, 20*blockHeight)
	ctx.Fill()

	ctx.SetSourceRGBA(0.5, 0.5, 0.5, 1)
	ctx.SetLineWidth(1)
	for i := 0; i < playfieldWidth; i++ {
		// Vertical lines
		ctx.MoveTo(float64(leftBorder+(i+1)*blockWidth), topBorder)
		ctx.LineTo(float64(leftBorder+(i+1)*blockWidth), topBorder+20*blockHeight)
		ctx.Stroke()
	}
	for i := 0; i < playfieldVisibleHeight; i++ {
		// Horizontal lines
		ctx.MoveTo(leftBorder, float64(topBorder+(i+1)*blockHeight))
		ctx.LineTo(leftBorder+10*blockWidth, float64(topBorder+(i+1)*blockHeight))
		ctx.Stroke()
	}
}

// drawFallenTetrominos : Draws the tetrominos the have already fallen to the "ground"
func (t *Tetris) drawFallenTetrominos(da *gtk.DrawingArea, ctx *cairo.Context) {
	for y := 0; y < playfieldVisibleHeight; y++ {
		for x := 0; x < playfieldVisibleWidth; x++ {
			idx := t.game.playfield[y][x]
			if idx > 0 {
				left, top := coordsToScreenCoords(x, y)
				t.drawBlock(da, ctx, tetrominos[idx-1].color, left, top)
			}
		}
	}
}

// drawFallingTetromino : Draws the currently falling tetromino
func (t *Tetris) drawFallingTetromino(da *gtk.DrawingArea, ctx *cairo.Context, tetro tetromino) {
	left, top := coordsToScreenCoords(t.game.falling.x, t.game.falling.y)

	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoHeight; x++ {
			if !tetro.blocks[y][x] {
				continue
			}
			if t.game.falling.y-y > playfieldVisibleHeight-1 {
				continue
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
