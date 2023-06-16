package tetris

import (
	"image/color"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

// onDraw : The onDraw signal handler
func (g *game) onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	g.drawBackground(da, ctx)
	g.drawField(da, ctx)
	g.drawFallenTetrominos(da, ctx)
	left, top := screenCoords(g.falling.x, g.falling.y)
	g.drawTetrominoAt(da, ctx, g.falling.tetro, left, top, false)
	g.drawQueuedTetrominos(da, ctx)
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

// drawField : Draws the playing field
func (g *game) drawField(_ *gtk.DrawingArea, ctx *cairo.Context) {
	// Draw the white background of the playing field
	ctx.SetSourceRGBA(1, 1, 1, 1)
	ctx.Rectangle(leftBorder, topBorder, 10*blockWidth, 22*blockHeight)
	ctx.Fill()
	// Draw the warning area of the field
	ctx.SetSourceRGBA(1, 0.5, 0, 0.5)
	ctx.Rectangle(leftBorder, topBorder, 10*blockWidth, 5*blockHeight)
	ctx.Fill()
	// Draw the danger area of the field
	ctx.SetSourceRGBA(1, 0, 0, 0.5)
	ctx.Rectangle(leftBorder, topBorder, 10*blockWidth, 2*blockHeight)
	ctx.Fill()

	// Draw grid lines
	ctx.SetSourceRGBA(0.5, 0.5, 0.5, 1)
	ctx.SetLineWidth(1)
	for i := 0; i < fieldWidth; i++ {
		// Vertical lines
		ctx.MoveTo(float64(leftBorder+(i+1)*blockWidth)+0.5, topBorder)
		ctx.LineTo(float64(leftBorder+(i+1)*blockWidth)+0.5, topBorder+fieldVisibleHeight*blockHeight)
		ctx.Stroke()
	}
	for i := 0; i < fieldVisibleHeight; i++ {
		// Horizontal lines
		ctx.MoveTo(leftBorder, float64(topBorder+(i+1)*blockHeight)+0.5)
		ctx.LineTo(leftBorder+10*blockWidth, float64(topBorder+(i+1)*blockHeight)+0.5)
		ctx.Stroke()
	}
}

// drawFallenTetrominos : Draws the tetrominos that have already fallen to the "ground"
func (g *game) drawFallenTetrominos(da *gtk.DrawingArea, ctx *cairo.Context) {
	for y := 0; y < fieldVisibleHeight; y++ {
		for x := 0; x < fieldVisibleWidth; x++ {
			idx := g.field[y][x]
			if idx > 0 {
				left, top := screenCoords(x, y)
				if y >= fieldLoosingHeight {
					// Game over
					g.quit()
				}
				g.drawBlock(da, ctx, tetrominos[idx-1].color, left, top)
			}
		}
	}
}

// drawTetrominoAt : Draws the currently falling tetromino at a certain position
func (g *game) drawTetrominoAt(da *gtk.DrawingArea, ctx *cairo.Context, tetro tetromino, left, top float64, queue bool) float64 {
	maxY := 0.0
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoHeight; x++ {
			if tetro.blocks[y][x] == 0 {
				continue
			}
			if !queue && g.falling.y-y > fieldVisibleHeight-1 {
				continue
			}
			maxY = top + float64(y)*blockHeight
			g.drawBlock(da, ctx, tetro.color, left+float64(x)*blockWidth, top+float64(y)*blockHeight)
		}
	}
	return maxY
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

// drawQueuedTetrominos : Draw the next 6 tetrominos
func (g *game) drawQueuedTetrominos(da *gtk.DrawingArea, ctx *cairo.Context) {
	q := g.rand.Queue()

	y := float64(topBorder)
	for i := 0; i < len(q); i++ {
		y = g.drawTetrominoAt(da, ctx, tetrominos[q[i]], 250, y, true) + blockHeight
	}

	ctx.SetSourceRGBA(1, 1, 1, 1)
	ctx.MoveTo(225, topBorder+65)
	ctx.LineTo(230, topBorder+22*blockHeight-10)
	ctx.MoveTo(235, topBorder+65)
	ctx.LineTo(230, topBorder+22*blockHeight-10)

	ctx.MoveTo(230, topBorder+60)
	ctx.LineTo(215, topBorder+75)
	ctx.MoveTo(230, topBorder+60)
	ctx.LineTo(245, topBorder+75)
	ctx.Stroke()

	ctx.SelectFontFace("Roboto Thin", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	ctx.SetFontSize(20)
	// ctx.TextExtents("Next")
	ctx.MoveTo(205, topBorder+45)
	ctx.ShowText("Next")
}
