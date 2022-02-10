package tetris

import (
	"math/rand"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

var falling tetromino
var posX, posY int
var quitChannel chan struct{}
var ticker *time.Ticker
var isPlaying = false
var playground [25][10]int

type Tetris struct {
	w  *gtk.ApplicationWindow
	da *gtk.DrawingArea
}

func NewTetris(w *gtk.ApplicationWindow, da *gtk.DrawingArea) *Tetris {
	t := &Tetris{w: w, da: da}
	t.w.Connect("key-press-event", t.onKeyPressed)
	t.da.Connect("draw", t.onDraw)
	return t
}

func (t *Tetris) StartGame() {
	isPlaying = true
	rand.Seed(time.Now().UnixNano())
	newFallingTetromino()

	ticker = time.NewTicker(250 * time.Millisecond)
	quitChannel = make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				posY -= 1
				t.da.QueueDraw()
				if t.checkBlockBottomSide() {
					newFallingTetromino()
				}
			case <-quitChannel:
				isPlaying = false
				ticker.Stop()
				return
			}
		}
	}()
}
