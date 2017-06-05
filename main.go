package main

import (
	"fmt"
	"math/rand"

	"bitbucket.org/cacilhas/gravity/system"
	"github.com/veandco/go-sdl2/sdl"
)

const spaceScale = 1e-10
const timeScale = 1e+9

func main() {
	window := initializeSDL(800, 800)
	defer window.Destroy()
	defer sdl.Quit()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	system := initializeSystem()

	for {
		plotSystem(surface, system)
		delay := rand.Float64()*90 + 10
		window.UpdateSurface()
		sdl.Delay(uint32(delay))
		system.Step(delay * timeScale / 1000)
	}
}

func plotSystem(surface *sdl.Surface, system gravity.System) {
	rect := sdl.Rect{X: 0, Y: 0, W: 800, H: 800}
	surface.FillRect(&rect, 0x00002255)
	center := system.GetBody("Sun").GetPosition()
	for _, body := range system.GetBodies() {
		plotBody(surface, body, center)
	}
}

func plotBody(surface *sdl.Surface, body gravity.Body, center gravity.Point) {
	pos := body.GetPosition().Add(center.Mul(-1))
	brect := sdl.Rect{
		X: int32(pos.GetX()*spaceScale) + 400,
		Y: int32(pos.GetY()*spaceScale) + 400,
		W: 2,
		H: 2,
	}
	surface.FillRect(&brect, 0x00ffffff)
}

func initializeSDL(width, height int) *sdl.Window {
	sdl.Init(sdl.INIT_EVERYTHING)
	window, err := sdl.CreateWindow(
		"Gravity",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		width, height,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}

	return window
}

func initializeSystem() gravity.System {
	body, _ := gravity.NewBody("Sun", 2e+30, 0, 0, 0)
	system, _ := gravity.NewSystem(body)

	for i := 1; i < 11; i++ {
		mass := 1.3e+22 + rand.Float64()*2e+27
		x := rand.Float64()*4.5e+9 - 2.25e+9
		y := rand.Float64()*4.5e+9 - 2.25e+9
		body, _ = gravity.NewBody(fmt.Sprintf("Planet %v", i), mass, x, y, 0)
		// TODO: fix inertia
		inertia := gravity.NewPoint(
			rand.Float64()*2.25e+9-x,
			rand.Float64()*2.25e+9-y,
			0,
		)
		body.SetInertia(inertia)
		system.AddBody(body)
	}

	return system
}
