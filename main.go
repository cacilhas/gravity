package main

import "github.com/veandco/go-sdl2/sdl"

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	window, err := sdl.CreateWindow(
		"Gravity",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		800, 800,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer sdl.Quit()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	// TODO: create and plot system

	rect := sdl.Rect{X: 0, Y: 0, W: 800, H: 800}
	surface.FillRect(&rect, 0x00002255)
	window.UpdateSurface()

	sdl.Delay(1000)
}
