package gui

import (
	"fmt"
	"github.com/davidnotplay/gameoflife/game"
	"github.com/gopherjs/gopherjs/js"
	"time"
)

const iconsContainerId string = "icons-container"
const inVarName string = "menuInterval"
const aEleSelector string = "#menu-container .animation"

// Transform an javascript object in javascript array
func objToArray(arr *js.Object) *js.Object {
	return js.Global.Get("Array").Call("from", arr)
}

func getById(id string) *js.Object {
	return js.Global.Get("document").Call("getElementById", id)
}

func showIcon(icon string) {
	container := getById(iconsContainerId)
	icons := container.Call("getElementsByTagName", "img")
	iconSel := container.Call("querySelector", fmt.Sprintf("#%s", icon))

	// Show and hide the images.
	objToArray(icons).Call("forEach", func(icon *js.Object) {
		icon.Set("style", "display:none")
		icon.Set("className", "")
	})
	container.Set("className", "show")
	iconSel.Set("style", "display:block")
	iconSel.Set("className", "start-animation")

	time.Sleep(400 * time.Millisecond)
	container.Set("className", "")
}

func handlerError(err error) {
	js.Global.Get("console").Call("error", err.Error())
	js.Global.Call("alert", "Error")
}


func getFuncWillShowGameInfo() func(c *Canvas) {
	msgEl := getById("menu-message")
	msg := "Size: %dx%d. Cells enabled %d. Cycles: %d."

	return func(c *Canvas) {
		m := c.game.GetMatrix()
		w, h, p := m.GetWidth(), m.GetHeight(), m.GetPointsEnabled()
		cn := c.game.GetCyclesNum()
		text := fmt.Sprintf(msg, w, h, p, cn)
		msgEl.Set("innerHTML", text)
	}
}


func togglePlayingGame(canvas *Canvas) {
	gameItemImg := getById("menu-game").Call("getElementsByTagName", "img").Index(0)

	if canvas.IsPlaying() {
		canvas.Stop()
		showIcon("stop-icon")
		println("Stoped")
		gameItemImg.Set("src", "./images/play-b.svg")

		return
	}

	// Start the game async
	go func() {
		if err := canvas.Start(getFuncWillShowGameInfo()); err != nil {
			canvas.Stop()
			handlerError(err)
		}
	}()

	showIcon("play-icon")
	println("Playing")
	gameItemImg.Set("src", "./images/stop-b.svg")
}

func showModal(modalId string) {
	modalc := getById("modal-container")
	modals := modalc.Call("querySelector", "#modal-container > div")
	modals = objToArray(modals)

	// hidden the modals.
	for i := 0; i < modals.Get("length").Int(); i++ {
		modals.Index(i).Set("style", "display: none")
	}


	// show the modal selected.
	getById(modalId).Set("style", "display: block")
	// make animatiion.
	modalc.Set("className", "show-modal")
}

func handleMenu(canvas *Canvas) {
	// set global variable will save the interval.
	js.Global.Set(inVarName, nil)

	// Get the elements animatables.
	aElems := objToArray(js.Global.Get("document").Call("querySelectorAll", aEleSelector))

	// func to changes the class in animatables elements.
	changeClass := func(cname string) {
		for i := 0; i < aElems.Get("length").Int(); i++ {
			aElems.Index(i).Set("className", cname)
		}
	}

	// Event will show the animate elements when the mouse is moved.
	canvas.GetJsCanvas().Call("addEventListener", "mousemove", func() { //  mousemove event.
		go func() {
			intId := js.Global.Get(inVarName)

			if intId != nil {
				js.Global.Call("clearTimeout", intId)
			} else {
				changeClass("animation animation-enter")
			}

			intId = js.Global.Call("setTimeout",
				func() {
					changeClass("animation animation-leave")
					js.Global.Set(inVarName, nil)
				},
				3000,
			)

			js.Global.Set(inVarName, intId)
		}()
	})

	// On-click info menu item
	getById("menu-info").Call("addEventListener", "click", func(evt *js.Object) {
		go func() {
			evt.Call("preventDefault")
			showModal("modal-info")
		}()
	})

	// on click game menu item.
	getById("menu-game").Call("addEventListener", "click", func(evt *js.Object) {
		go func() {
			evt.Call("preventDefault")
			togglePlayingGame(canvas)
		}()
	})
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
			canvas, _ = NewCanvas(&[]game.Position{})
		}()
	})

	handleMenu(canvas)

	// Close modal function.
	closeModalFun := func() {
		go func() {
			getById("modal-container").Set("className", "hide-modal")
		}()
	}

	// Close modal when it clicks in the button.
	getById("close-modal").Call("addEventListener", "click", closeModalFun)
	// Close modal when it clicks in the modal overlay.
	getById("modal-container").Call("addEventListener", "click", closeModalFun)
}
