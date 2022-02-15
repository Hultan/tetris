package tetris

import (
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/tetris/internal/randomizer"
)

type Tetris struct {
	window      *gtk.ApplicationWindow
	drawingArea *gtk.DrawingArea

	game *game
}

func NewTetris(w *gtk.ApplicationWindow, da *gtk.DrawingArea) *Tetris {
	t := &Tetris{window: w, drawingArea: da}
	t.window.Connect("key-press-event", t.onKeyPressed)
	return t
}

func (t *Tetris) StartGame() {
	t.game = &game{}
	// TODO : Move to game constructor
	t.game.isActive = true
	t.game.rand = randomizer.NewRandomizer(tetrominoCount, queueSize)
	t.game.createNewFallingTetromino()
	t.game.speed = 500
	t.drawingArea.Connect("draw", t.game.onDraw)

	t.game.ticker.ticker = time.NewTicker(t.game.speed * time.Millisecond)
	t.game.ticker.tickerQuit = make(chan struct{})

	go t.mainLoop()
}

func (t *Tetris) mainLoop() {
	for {
		select {
		case <-t.game.ticker.ticker.C:
			t.drawingArea.QueueDraw()
			if t.game.checkPlayfieldBottom() {
				t.game.createNewFallingTetromino()
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
	case 97: // Button "A" => Move tetromino left
		if t.game.isActive && !t.game.checkPlayfieldSides(true) {
			t.game.falling.x -= 1
		}
	case 113: // Button "Q" => Quit game
		t.game.quit()
		t.window.Close() // Close window
	case 115: // Button "S" => Rotate tetromino
		if t.game.isActive {
			t.game.rotateTetromino(&t.game.falling.tetro)
		}
	case 100: // Button "D" => Move tetromino right
		if t.game.isActive && !t.game.checkPlayfieldSides(false) {
			t.game.falling.x += 1
		}
	case 120: // Button "X" => Move tetromino down
		if t.game.isActive {
			t.game.dropTetrominoToPlayfield()
			t.game.createNewFallingTetromino()
		}
	}
	t.drawingArea.QueueDraw()
}

func (g *game) quit() {
	if g.isActive {
		g.isActive = false
		close(g.ticker.tickerQuit) // Stop ticker
	}
}
