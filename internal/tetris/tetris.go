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

	isPlaying  bool
	playground [25][10]int
	current    tetromino
	posX, posY int
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
	t.isPlaying = true
	rand.Seed(time.Now().UnixNano())
	t.createNewFallingTetromino()

	t.ticker = time.NewTicker(500 * time.Millisecond)
	t.tickerQuit = make(chan struct{})
	go func() {
		for {
			select {
			case <-t.ticker.C:
				t.posY -= 1
				t.da.QueueDraw()
				if t.checkBlockBottomSide() {
					t.createNewFallingTetromino()
				}
			case <-t.tickerQuit:
				t.isPlaying = false
				t.ticker.Stop()
				return
			}
		}
	}()
}

func (t *Tetris) adjustPositionAfterRotate() {
	min, max := 0, 9
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !t.current.blocks[y][x] {
				continue
			}
			if t.posX+x < min {
				min = t.posX + x
			}
			if t.posX+x > max {
				max = t.posX + x
			}
		}
	}

	if min < 0 {
		t.posX += -min
	}
	if max > 9 {
		t.posX -= max - 9
	}
}

// Drop a new Tetromino
func (t *Tetris) createNewFallingTetromino() {
	r := rand.Intn(7)
	t.current = tetrominos[r]
	t.posY = 24
	t.posX = 3
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
