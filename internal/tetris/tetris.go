package tetris

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

type Tetris struct {
	w  *gtk.ApplicationWindow
	da *gtk.DrawingArea

	game game

	isActive   bool
	playground [playgroundHeight][playgroundWidth]int
	falling    fallingTetromino
	ticker     ticker
}

type game struct {
}

type fallingTetromino struct {
	tetro tetromino
	x, y  int
}

type ticker struct {
	tickerQuit chan struct{}
	ticker     *time.Ticker
}

func NewTetris(w *gtk.ApplicationWindow, da *gtk.DrawingArea) *Tetris {
	t := &Tetris{w: w, da: da}
	t.w.Connect("key-press-event", t.onKeyPressed)
	t.da.Connect("draw", t.onDraw)
	return t
}

func (t *Tetris) StartGame() {
	t.isActive = true
	rand.Seed(time.Now().UnixNano())
	t.createNewFallingTetromino()

	t.ticker.ticker = time.NewTicker(500 * time.Millisecond)
	t.ticker.tickerQuit = make(chan struct{})
	go func() {
		for {
			select {
			case <-t.ticker.ticker.C:
				t.falling.y -= 1
				t.da.QueueDraw()
				if t.checkBlockBottomSide() {
					t.createNewFallingTetromino()
				}
			case <-t.ticker.tickerQuit:
				t.isActive = false
				t.ticker.ticker.Stop()
				return
			}
		}
	}()
}

func (t *Tetris) quitGame() {
	if t.isActive {
		close(t.ticker.tickerQuit) // Stop ticker
	}
	t.w.Close() // Close window
}

func (t *Tetris) adjustPositionAfterRotate() {
	min, max := 0, playgroundWidth-1
	for y := 0; y < tetrominoHeight; y++ {
		for x := 0; x < tetrominoWidth; x++ {
			if !t.falling.tetro.blocks[y][x] {
				continue
			}
			if t.falling.x+x < min {
				min = t.falling.x + x
			}
			if t.falling.x+x > max {
				max = t.falling.x + x
			}
		}
	}

	if min < 0 {
		t.falling.x += -min
	}
	if max > playgroundWidth-1 {
		t.falling.x -= max - (playgroundWidth - 1)
	}
}

// Drop a new Tetromino
func (t *Tetris) createNewFallingTetromino() {
	r := rand.Intn(tetrominoCount)
	t.falling.tetro = tetrominos[r]
	t.falling.y = playgroundHeight - 1
	t.falling.x = (playgroundWidth - tetrominoWidth) / 2
}

// Rotate the 4x4 tetromino array 90 degrees
// https://www.geeksforgeeks.org/rotate-a-matrix-by-90-degree-in-clockwise-direction-without-using-any-extra-space/
func (t *Tetris) rotateTetromino(tetro *tetromino) {
	for y := 0; y < tetrominoHeight/2; y++ {
		for x := y; x < tetrominoWidth-y-1; x++ {
			tmp := tetro.blocks[y][x]
			tetro.blocks[y][x] = tetro.blocks[tetrominoHeight-1-x][y]
			tetro.blocks[tetrominoHeight-1-x][y] = tetro.blocks[tetrominoHeight-1-y][tetrominoWidth-1-x]
			tetro.blocks[tetrominoHeight-1-y][tetrominoWidth-1-x] = tetro.blocks[x][tetrominoWidth-1-y]
			tetro.blocks[x][tetrominoWidth-1-y] = tmp
		}
	}

	t.adjustPositionAfterRotate()
}

// Debug function : Print the playground
func (t *Tetris) printPlayground() {
	fmt.Println()
	fmt.Println("----------------")
	fmt.Println()
	for r := 0; r < playgroundHeight; r++ {
		fmt.Println(t.playground[r])
	}
}
