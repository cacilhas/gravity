package gravity

import (
	"fmt"
	"math"
)

// Body represents an object on a system
type Body interface {
	GetName() string
	GetMass() float64
	GetPosition() Point
	SetPosition(Point)
	GetInertia() Point
	SetInertia(v Point)
	Move(float64) error
	Grav(Body) Point
	String() string
}

type body struct {
	name     string
	mass     float64
	position Point
	inertia  Point
}

// NewBody create a new Body
func NewBody(name string, mass, x, y, z float64) (Body, error) {
	if mass <= 0 {
		return nil, fmt.Errorf("invalid mass: %v", mass)
	}

	obj := body{
		name:     name,
		mass:     mass,
		position: NewPoint(x, y, z),
		inertia:  NewPoint(0, 0, 0),
	}
	return &obj, nil
}

func (b body) GetName() string {
	return b.name
}

func (b body) GetMass() float64 {
	return b.mass
}

func (b body) GetPosition() Point {
	return b.position
}

func (b *body) SetPosition(pos Point) {
	b.position = pos
}

func (b body) GetInertia() Point {
	return b.inertia
}

func (b *body) SetInertia(v Point) {
	b.inertia = v
}

func (b *body) Move(dt float64) error {
	if dt < 0 {
		return fmt.Errorf("invalid timedelta %v", dt)
	}

	mass := b.GetMass()

	if mass <= 0 {
		return fmt.Errorf("invalid mass %v", mass)
	}

	inertia := b.GetInertia()
	movement := NewPoint(
		inertia.GetX()*dt/mass,
		inertia.GetY()*dt/mass,
		inertia.GetZ()*dt/mass,
	)
	b.position = b.position.Add(movement)
	return nil
}

func (b body) Grav(other Body) Point {
	diff := b.position.Diff(other.GetPosition())
	d := diff.Magnitude()
	f := G * b.mass * other.GetMass() / (d * d)

	dx := diff.GetX()
	if dx != 0 {
		dx /= math.Abs(dx)
		dx *= -f
	}

	dy := diff.GetY()
	if dy != 0 {
		dy /= math.Abs(dy)
		dy *= -f
	}

	dz := diff.GetZ()
	if dz != 0 {
		dz /= math.Abs(dz)
		dz *= -f
	}

	return NewPoint(dx, dy, dz)
}

func (b body) String() string {
	return fmt.Sprintf(
		"{%v: %vKg @%v F%v}", b.name, b.mass, b.position, b.inertia,
	)
}
