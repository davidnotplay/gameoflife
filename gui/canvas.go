package gui

import (
	"time"
	"math"
	"github.com/gopherjs/gopherjs/js"
	"github.com/davidnotplay/gameoflife/game"
	"github.com/davidnotplay/gameoflife/matrix"
)

// Constant pixels per points.
const ppp int = 14

// Canvas color cells whe these are enabled or disabled.
var pointColors map[bool]string = map[bool]string{false: "#444", true: "#ff0"}

type Canvas struct {
	// Js canvas object.
	canvas *js.Object

	// Game matrix.
	game *game.Game

	playing bool
}

// Get the canvas js object from html dom.
func getCanvas() *js.Object {
	return js.Global.Get("document").Call("getElementById", "game-of-life")
}

// Get the game size (width, height) depending of the window browser size.
func getGameSize() (int, int){
	canvas := js.Global.Get("document").Call("getElementById", "game-of-life")

	// Set the canvas size using the windows size.
	ww := js.Global.Get("window").Get("innerWidth").Float()
	wh := js.Global.Get("window").Get("innerHeight").Float()
	canvas.Set("width", ww)
	canvas.Set("height", wh)

	// Generate the matrix size using the window size.
	gw := int(math.Ceil(ww/float64(ppp)))
	gh := int(math.Ceil(wh/float64(ppp)))

	return gw, gh
}

// Make the canvas using the matrix data.
func (self *Canvas)generate() error {
	matrix := self.game.GetMatrix()
	w, h := matrix.GetWidth(), matrix.GetHeight()
	ctx := self.canvas.Call("getContext", "2d")

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			enabled, err := matrix.IsEnabled(i, j)
			if err != nil {
				return err
			}

			ctx.Set("fillStyle", pointColors[enabled])
			ctx.Call("fillRect", i*ppp+1, j*ppp+1, ppp - 1, ppp - 1)
		}
	}

	return nil
}

// Make new Canvas for the game.
// The function make the html5 canvas and prepare the game using the initial positions `p`
func NewCanvas(p *[]game.Position) (*Canvas, error) {
	gw, gh := getGameSize()
	game, err := game.New(gw, gh, *p)

	if err != nil {
		// Error creating the game.
		return nil, err
	}

	canvas := &Canvas{getCanvas(), game, false}
	err = canvas.generate()

	if err != nil {
		// Error generating the canvas.
		return nil, err
	}

	return canvas, nil
}

func (self *Canvas)run() error {
	var err error

	for ;self.playing; {
		time.Sleep(200 * time.Millisecond)
		err = self.game.Cycle()

		if err != nil {
			return err
		}

		err = self.generate()

		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Canvas)Start() error {
	self.playing = true
	return self.run()
}

func (self *Canvas)Stop() {
	self.playing = false
}

// Returns the js object where are saved the canvas.
func (self *Canvas)GetJsCanvas() *js.Object {
	return self.canvas
}

func (self *Canvas)ToggleMatrixPoint(x, y int) error {
	var err error = nil
	mx := x / ppp
	my := y / ppp
	m := self.game.GetMatrix()
	enabled, err := m.IsEnabled(mx, my)

	if err != nil {
		return err
	}

	if enabled {
		err = m.DisablePoint(mx, my)
	} else {
		err = m.EnablePoint(mx, my)
	}

	if err != nil {
		return err
	}

	// Redraw canvas.
	err = self.generate()
	return err
}

func (self *Canvas)IsPlaying() bool {
	return self.playing
}

