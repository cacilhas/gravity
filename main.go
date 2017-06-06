package main

import (
	"fmt"
	"math/rand"

	"time"

	"bitbucket.org/cacilhas/gravity/system"
	"github.com/veandco/go-sdl2/sdl"
)

const timeScale = 1e+12
const wsize = 600
const wdiag = wsize / 2

var spaceScale float64

func main() {
	window := initializeSDL(wsize, wsize)
	defer window.Destroy()
	defer sdl.Quit()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	rand.Seed(int64(time.Now().Nanosecond()))
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

	var futher float64
	for _, body := range system.GetBodies() {
		pos := body.GetPosition()
		if x := pos.GetX(); x > futher {
			futher = x
		}
		if y := pos.GetY(); y > futher {
			futher = y
		}
	}
	spaceScale = wdiag / futher

	for _, body := range system.GetBodies() {
		plotBody(surface, body, center)
	}
}

func plotBody(surface *sdl.Surface, body gravity.Body, center gravity.Point) {
	pos := body.GetPosition().Add(center.Mul(-1))
	brect := sdl.Rect{
		X: int32(pos.GetX()*spaceScale) + wdiag,
		Y: int32(pos.GetY()*spaceScale) + wdiag,
		W: 2,
		H: 2,
	}
	surface.FillRect(&brect, 0x00ffffff)
}

func initializeSDL(width, height int) *sdl.Window {
	sdl.Init(sdl.INIT_EVERYTHING)
	window, err := sdl.CreateWindow(
		"Gravity",
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
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
	centerMass := body.GetMass()

	for i := 1; i <= 10; i++ {
		mass := 1.3e+22 + rand.Float64()*2e+27
		x := rand.Float64()*4.5e+9 - 2.25e+9
		y := rand.Float64()*4.5e+9 - 2.25e+9
		body, _ = gravity.NewBody(fmt.Sprintf("Planet %v", i), mass, x, y, 0)
		inertia := body.GetPosition().Mul(body.GetMass() * centerMass).TanXY()
		body.SetInertia(inertia)
		system.AddBody(body)
	}

	return system
}
