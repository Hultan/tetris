package tetris

import (
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/tetris/internal/randomizer"
)

type Tetris struct {
	window *gtk.ApplicationWindow
	da     *gtk.DrawingArea

	game *game
}

func NewTetris(w *gtk.ApplicationWindow, da *gtk.DrawingArea) *Tetris {
	t := &Tetris{window: w, da: da}
	t.window.Connect("key-press-event", t.onKeyPressed)
	return t
}

func (t *Tetris) StartGame() {
	t.game = &game{}
	// TODO : Move to game constructor
	t.game.isActive = true
	t.game.rand = randomizer.NewRandomizer(tetrominoCount, queueSize)
	t.game.nextTetromino()
	t.game.speed = 500
	t.da.Connect("draw", t.game.onDraw)

	t.game.ticker.ticker = time.NewTicker(t.game.speed * time.Millisecond)
	t.game.ticker.tickerQuit = make(chan struct{})

	go t.mainLoop()
}

func (t *Tetris) mainLoop() {
	for {
		select {
		case <-t.game.ticker.ticker.C:
			t.da.QueueDraw()
			if t.game.checkFieldBottom() {
				t.game.nextTetromino()
			}
			t.game.falling.y -= 1
		case <-t.game.ticker.tickerQuit:
			t.game.isActive = false
			t.game.ticker.ticker.Stop()
			return
		}
	}
}

// onKeyPressed : The onKeyPressed signal handler
func (t *Tetris) onKeyPressed(_ *gtk.ApplicationWindow, e *gdk.Event) {
	key := gdk.EventKeyNewFromEvent(e)

	switch key.KeyVal() {
	case gdk.KEY_Q, gdk.KEY_q: // Button "Q" => Quit game
		t.game.quit()
		t.window.Close() // Close window
	case gdk.KEY_A, gdk.KEY_a, gdk.KEY_Left: // Button "A" => Move tetromino left
		if t.game.isActive && !t.game.checkFieldSides(true) {
			t.game.falling.x -= 1
		}
	case gdk.KEY_W, gdk.KEY_w, gdk.KEY_Up: // Button "W" => Rotate tetromino
		if t.game.isActive {
			t.game.rotateTetromino(&t.game.falling.tetro)
		}
	case gdk.KEY_D, gdk.KEY_d, gdk.KEY_Right: // Button "D" => Move tetromino right
		if t.game.isActive && !t.game.checkFieldSides(false) {
			t.game.falling.x += 1
		}
	case gdk.KEY_S, gdk.KEY_s, gdk.KEY_Down: // Button "S" => Move tetromino down
		if t.game.isActive {
			t.game.dropTetrominoToField()
			t.game.nextTetromino()
		}
	}
	t.da.QueueDraw()
}

func (g *game) quit() {
	if g.isActive {
		g.isActive = false
		close(g.ticker.tickerQuit) // Stop ticker
	}
}
