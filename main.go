package main

import (
	"fmt"
	"math"
	"math/rand"
	"path/filepath"
	"sync"
	"time"

	"bitbucket.org/cacilhas/gravity/system"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

const timeScale = 1e+12
const wsize = 600
const wradius = wsize / 2

var spaceScale, tick float64
var sphere *sdl.Surface

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
		window.UpdateSurface()
		// TODO: make it work
		if state := sdl.GetKeyboardState(); state[sdl.SCANCODE_ESCAPE] != 0 {
			break
		}
		count := len(system.GetBodies())
		fmt.Printf(
			"bodies: %2d\tscale: %f         \r",
			count, -math.Log10(spaceScale),
		)
		wait(system)
	}
}

func wait(system gravity.System) {
	sdl.Delay(100)
	now := float64(time.Now().UnixNano()) * 10e-6
	if tick != 0 {
		system.Step(now - tick)
	}
	tick = now
}

func plotSystem(surface *sdl.Surface, system gravity.System) {
	surface.FillRect( // background
		&sdl.Rect{X: 0, Y: 0, W: wsize, H: wsize},
		0x00002255,
	)
	center := system.GetBody("Sun").GetPosition()
	bodies := system.GetBodies()

	var futher float64
	for _, body := range bodies {
		pos := body.GetPosition()
		futher = math.Max(
			futher,
			math.Max(math.Abs(pos.GetX()), math.Abs(pos.GetY())),
		)
	}
	spaceScale = math.Max(wradius/futher, 1e-8)
	bodies = system.GetBodies()

	var lock sync.WaitGroup
	lock.Add(len(bodies))
	for _, body := range bodies {
		go plotBody(surface, body, center, &lock)
	}
	lock.Wait()
}

func plotBody(surface *sdl.Surface, body gravity.Body, center gravity.Point, lock *sync.WaitGroup) {
	defer lock.Done()
	rect := calculatePosition(body, center)

	if rect.W == 0 { // just a dot
		rect.W = 1
		rect.H = 1
		surface.FillRect(rect, 0x00ffffff)

	} else { // big enough
		sphere.BlitScaled(
			&sdl.Rect{X: 0, Y: 0, W: sphere.W, H: sphere.H},
			surface,
			rect,
		)
	}
}

func calculatePosition(body gravity.Body, center gravity.Point) *sdl.Rect {
	pos := body.GetPosition().Add(center.Mul(-1))
	radius := int32(body.GetMass() / 2e+29)

	rect := sdl.Rect{
		X: int32(pos.GetX()*spaceScale) + wradius,
		Y: int32(pos.GetY()*spaceScale) + wradius,
		W: radius,
		H: radius,
	}
	return &rect
}

func initializeSDL(width, height int) *sdl.Window {
	sdl.Init(sdl.INIT_EVERYTHING)
	img.Init(img.INIT_PNG)
	window, err := sdl.CreateWindow(
		"Gravity",
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		width, height,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}
	// TODO: this stinks
	filename, _ := filepath.Abs("./sphere.png")
	sphere, err = img.Load(filename)
	if err != nil {
		panic(err)
	}

	return window
}

func initializeSystem() gravity.System {
	body, _ := gravity.NewBody("Sun", 2e+30, 0, 0, 0)
	system, _ := gravity.NewSystem(body)

	for i := 1; i <= 10; i++ {
		mass := 1.3e+22 + rand.Float64()*2e+27
		x := rand.Float64()*4.5e+9 - 2.25e+9
		y := rand.Float64()*4.5e+9 - 2.25e+9
		body, _ = gravity.NewBody(fmt.Sprintf("Planet %v", i), mass, x, y, 0)
		inertia := body.GetPosition().TanXY().Mul(5e+21)
		body.SetInertia(inertia)
		system.AddBody(body)
	}

	return system
}
