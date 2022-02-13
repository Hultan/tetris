package tetris

import (
	"math/rand"
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Tetris struct {
	window      *gtk.ApplicationWindow
	drawingArea *gtk.DrawingArea

	game *game
}

type game struct {
	speed     time.Duration
	isActive  bool
	playfield [playfieldHeight][playfieldWidth]int
	falling   fallingTetromino
	ticker    ticker
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
	t := &Tetris{window: w, drawingArea: da}
	t.window.Connect("key-press-event", t.onKeyPressed)
	t.drawingArea.Connect("draw", t.onDraw)
	return t
}

func (t *Tetris) StartGame() {
	t.game = &game{}
	t.game.isActive = true
	rand.Seed(time.Now().UnixNano())
	t.createNewFallingTetromino()
	t.game.speed = 500

	t.game.ticker.ticker = time.NewTicker(t.game.speed * time.Millisecond)
	t.game.ticker.tickerQuit = make(chan struct{})
	go func() {
		for {
			select {
			case <-t.game.ticker.ticker.C:
				t.game.falling.y -= 1
				t.drawingArea.QueueDraw()
				if t.checkPlayfieldBottom() {
					t.createNewFallingTetromino()
				}
			case <-t.game.ticker.tickerQuit:
				t.game.isActive = false
				t.game.ticker.ticker.Stop()
				return
			}
		}
	}()
}

func (t *Tetris) quitGame() {
	if t.game.isActive {
		t.game.isActive = false
		close(t.game.ticker.tickerQuit) // Stop ticker
	}
	t.window.Close() // Close window
}

// onKeyPressed : The onKeyPressed signal handler
func (t *Tetris) onKeyPressed(_ *gtk.ApplicationWindow, e *gdk.Event) {
	key := gdk.EventKeyNewFromEvent(e)

	switch key.KeyVal() {
	case 97: // Button "A" => Move tetromino left
		if !t.checkPlayfieldSides(true) {
			t.game.falling.x -= 1
		}
	case 113: // Button "Q" => Quit game
		t.quitGame()
	case 115: // Button "S" => Rotate tetromino
		t.rotateTetromino(&t.game.falling.tetro)
	case 100: // Button "D" => Move tetromino right
		if !t.checkPlayfieldSides(false) {
			t.game.falling.x += 1
		}
	case 120: // Button "X" => Move tetromino down
		t.dropTetrominoToPlayfield()
	}
	t.drawingArea.QueueDraw()
}
