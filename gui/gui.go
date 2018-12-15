package gui

import (
	"fmt"
	"time"
	"github.com/gopherjs/gopherjs/js"
	"github.com/davidnotplay/gameoflife/game"
)

const iconsContainerId string = "icons-container"

func showIcon(icon string) {
	container := js.Global.Get("document").Call("getElementById", iconsContainerId)
	icons := container.Call("getElementsByTagName", "img")
	iconSel := container.Call("querySelector", fmt.Sprintf("#%s", icon))

	// Show and hide the images.
	js.Global.Get("Array").Call("from", icons).Call("forEach", func(icon *js.Object) {
		icon.Set("style", "display:none")
	})

	iconSel.Set("style", "display:block")


	container.Set("className", "icons-animation-start")
	// /** @TODO mejorar animacion. */
	time.Sleep(10 * time.Millisecond) // add time sleep to make the css animation.

	container.Set("className", "icons-animation-start icons-animation-end")
	time.Sleep(400 * time.Millisecond)
	container.Set("className", "")
}

func handlerError(err error) {
	js.Global.Get("console").Call("error", err.Error())
	js.Global.Call("alert", "Error")
}

func togglePlayingGame(canvas *Canvas) {
	if canvas.IsPlaying() {
		canvas.Stop()
		showIcon("stop-icon")
		println("Stoped")

		return
	}

	// Start the game async
	go func() {
		if err := canvas.Start(); err != nil {
			canvas.Stop()
			handlerError(err)
		}
	}()

	showIcon("play-icon")
	println("Playing")
}


func Start() {
	var canvas *Canvas
	canvas, _ = NewCanvas(&[]game.Position{})

	// click event in canvas. Enable or disable matrix points.
	canvas.GetJsCanvas().Call("addEventListener", "click", func(evt *js.Object) {
		go func() {
			x := evt.Get("clientX").Int()
			y := evt.Get("clientY").Int()
			canvas.ToggleMatrixPoint(x, y)
		}()
	})

	// Keypress event in window. Start or stop the game.
	js.Global.Get("window").Call("addEventListener", "keypress", func(evt *js.Object) {
		go func() {
			if evt.Get("charCode").Int() == int(' ') {
				togglePlayingGame(canvas)
			}
		}()
	})

	// Remake the canvas when the window size changes.
	// It is rare!!! The matrix position no disappers when the matrix is removed.
	// Maybe guilty are async functions.
	js.Global.Get("window").Call("addEventListener", "resize", func(evt *js.Object) {
		go func() {
			canvas.Stop()
			canvas, _ = NewCanvas(&[]game.Position{});
		}()
	})
}
