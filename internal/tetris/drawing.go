package tetris

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

// onDraw : The onDraw signal handler
func (g *game) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	g.drawBackground(da, ctx)
	g.drawPlayfield(da, ctx)
	g.drawFallenTetrominos(da, ctx)
	g.drawFallingTetromino(da, ctx, g.falling.tetro)
}

//
// HELPER FUNCTIONS
//

// drawBackground : Draws the background
func (g *game) drawBackground(da *gtk.DrawingArea, ctx *cairo.Context) {
	width := float64(da.GetAllocatedWidth())
	height := float64(da.GetAllocatedHeight())
	ctx.SetSourceRGB(0.4, 0.4, 1)
	ctx.Rectangle(0, 0, width, height)
	ctx.Fill()
}

// drawPlayfield : Draws the playfield background
func (g *game) drawPlayfield(da *gtk.DrawingArea, ctx *cairo.Context) {
	ctx.SetSourceRGBA(1, 1, 1, 1)
	ctx.Rectangle(leftBorder, topBorder, 10*blockWidth, 22*blockHeight)
	ctx.Fill()
	ctx.SetSourceRGBA(1, 0.5, 0, 0.5)
	ctx.Rectangle(leftBorder, topBorder, 10*blockWidth, 5*blockHeight)
	ctx.Fill()
	ctx.SetSourceRGBA(1, 0, 0, 0.5)
	ctx.Rectangle(leftBorder, topBorder, 10*blockWidth, 2*blockHeight)
	ctx.Fill()

	ctx.SetSourceRGBA(0.5, 0.5, 0.5, 1)
	ctx.SetLineWidth(1)
	for i := 0; i < playfieldWidth; i++ {
		// Vertical lines
		ctx.MoveTo(float64(leftBorder+(i+1)*blockWidth)+0.5, topBorder)
		ctx.LineTo(float64(leftBorder+(i+1)*blockWidth)+0.5, topBorder+playfieldVisibleHeight*blockHeight)
		ctx.Stroke()
	}
	for i := 0; i < playfieldVisibleHeight; i++ {
		// Horizontal lines
		ctx.MoveTo(leftBorder, float64(topBorder+(i+1)*blockHeight)+0.5)
		ctx.LineTo(leftBorder+10*blockWidth, float64(topBorder+(i+1)*blockHeight)+0.5)
		ctx.Stroke()
	}
}

// drawFallenTetrominos : Draws the tetrominos the have already fallen to the "ground"
func (g *game) drawFallenTetrominos(da *gtk.DrawingArea, ctx *cairo.Context) {
	for y := 0; y < playfieldVisibleHeight; y++ {
		for x := 0; x < playfieldVisibleWidth; x++ {
			idx := g.playfield[y][x]
			if idx > 0 {
				left, top := coordsToScreenCoords(x, y)
				if y >= playfieldLoosingHeight {
					// Game over
					g.quit()
				}
				g.drawBlock(da, ctx, tetrominos[idx-1].color, left, top)
			}
		}
	}
}

// drawFallingTetromino : Draws the currently falling tetromino
func (g *game) drawFallingTetromino(da *gtk.DrawingArea, ctx *cairo.Context, tetro tetromino) {
	left, top := coordsToScreenCoords(g.falling.x, g.falling.y)

	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoHeight; x++ {
			if tetro.blocks[y][x] == 0 {
				continue
			}
			if g.falling.y-y > playfieldVisibleHeight-1 {
				continue
			}
			g.drawBlock(da, ctx, tetro.color, left+float64(x)*blockWidth, top+float64(y)*blockHeight)
		}
	}
}

// drawBlock : Draws a single block
func (g *game) drawBlock(_ *gtk.DrawingArea, ctx *cairo.Context, c color.Color, left, top float64) {
	// Fill block in the correct color
	red, green, blue, alpha := c.RGBA()
	ctx.SetSourceRGBA(col(red), col(green), col(blue), col(alpha))
	ctx.Rectangle(left, top, blockWidth, blockHeight)
	ctx.Fill()

	// Draw black border around block
	ctx.SetSourceRGBA(0, 0, 0, 1)
	ctx.SetLineWidth(1)
	ctx.Rectangle(left+0.5, top+0.5, blockWidth, blockHeight)
	ctx.Stroke()
}
