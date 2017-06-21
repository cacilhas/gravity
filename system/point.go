package gravity

import (
	"fmt"
	"math"
)

// Point represents a 3D point
type Point interface {
	GetX() float64
	GetY() float64
	GetZ() float64
	Hypot() float64
	TanXY() Point
	TanXZ() Point
	TanYZ() Point

	Add(Point) Point
	Add2(x, y float64) Point
	Add3(x, y, z float64) Point
	Diff(Point) Point
	Mul(float64) Point

	String() string
}

// NewPoint creates a new point
func NewPoint(x, y, z float64) Point {
	return point{x, y, z}
}

type point struct {
	x, y, z float64
}

func (p point) GetX() float64 {
	return p.x
}

func (p point) GetY() float64 {
	return p.y
}

func (p point) GetZ() float64 {
	return p.z
}

func (p point) Hypot() float64 {
	return math.Hypot(math.Hypot(p.x, p.y), p.z)
}

func (p point) TanXY() Point {
	a := p.Hypot()
	ang := math.Acos(p.x/a) + (math.Pi / 2)
	x := a * math.Cos(ang)
	y := a * math.Sin(ang)
	return point{
		x: x,
		y: y,
		z: p.z,
	}
}

func (p point) TanXZ() Point {
	a := p.Hypot()
	ang := math.Acos(p.x/a) + (math.Pi / 2)
	x := a * math.Cos(ang)
	z := a * math.Sin(ang)
	return point{
		x: x,
		y: p.y,
		z: z,
	}
}

func (p point) TanYZ() Point {
	a := p.Hypot()
	ang := math.Acos(p.y/a) + (math.Pi / 2)
	y := a * math.Cos(ang)
	z := a * math.Sin(ang)
	return point{
		x: p.x,
		y: y,
		z: z,
	}
}

func (p point) Add(other Point) Point {
	return point{
		x: p.x + other.GetX(),
		y: p.y + other.GetY(),
		z: p.z + other.GetZ(),
	}
}

func (p point) Add2(x, y float64) Point {
	return point{
		x: p.x + x,
		y: p.y + y,
		z: p.z,
	}
}

func (p point) Add3(x, y, z float64) Point {
	return point{
		x: p.x + x,
		y: p.y + y,
		z: p.z + z,
	}
}

func (p point) Diff(other Point) Point {
	return point{
		x: p.x - other.GetX(),
		y: p.y - other.GetY(),
		z: p.z - other.GetZ(),
	}
}

func (p point) Mul(f float64) Point {
	return point{
		x: p.x * f,
		y: p.y * f,
		z: p.z * f,
	}
}

func (p point) String() string {
	return fmt.Sprintf("(%v, %v, %v)", p.x, p.y, p.z)
}
