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
	playground [25][10]int
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
	min, max := 0, 9
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
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
	if max > 9 {
		t.falling.x -= max - 9
	}
}

// Drop a new Tetromino
func (t *Tetris) createNewFallingTetromino() {
	r := rand.Intn(7)
	t.falling.tetro = tetrominos[r]
	t.falling.y = 24
	t.falling.x = 3
}

// Rotate the 5x5 tetromino array 90 degrees
func (t *Tetris) rotateTetromin(tetro *tetromino) {
	for i := 0; i < 5/2; i++ {
		for j := 0; j < 5-i-1; j++ {
			tmp := tetro.blocks[i][j]
			tetro.blocks[i][j] = tetro.blocks[5-1-j][i]
			tetro.blocks[5-1-j][i] = tetro.blocks[5-1-i][5-1-j]
			tetro.blocks[5-1-i][5-1-j] = tetro.blocks[j][5-1-i]
			tetro.blocks[j][5-1-i] = tmp
		}
	}

	t.adjustPositionAfterRotate()
}

// Debug function : Print the playground
func (t *Tetris) printPlayground() {
	fmt.Println()
	fmt.Println("----------------")
	fmt.Println()
	for r := 0; r < len(t.playground); r++ {
		fmt.Println(t.playground[r])
	}
}
