package gravity

import "fmt"

// G universal gravitational constant
const G = 6.67408e-11

// System represents a gravitational system
type System interface {
	Status() error
	GetBodies() map[string]Body
	GetBody(string) Body
	AddBody(Body) error
	RemoveBody(Body) bool
	Step(float64) error
	TotalMass() float64
	String() string
}

type system struct {
	status error
	bodies map[string]Body
}

// NewSystem build a new system
func NewSystem(args ...Body) (System, error) {
	s := system{nil, make(map[string]Body)}
	for _, b := range args {
		if name := b.GetName(); s.bodies[name] == nil {
			s.bodies[name] = b
		} else {
			return nil, fmt.Errorf("duplicated body: %v", name)
		}
	}
	return &s, nil
}

func (s system) Status() error {
	return s.status
}

func (s *system) GetBodies() map[string]Body {
	return s.bodies
}

func (s system) GetBody(name string) Body {
	return s.bodies[name]
}

func (s *system) AddBody(b Body) error {
	name := b.GetName()
	if s.bodies[name] == nil {
		s.bodies[name] = b
		return nil
	}

	return fmt.Errorf("body %v already exists", name)
}

func (s *system) RemoveBody(b Body) bool {
	name := b.GetName()
	if s.bodies[name] != b {
		return false
	}

	delete(s.bodies, name)
	return true
}

func (s *system) Step(dt float64) error {
	if s.status != nil {
		return s.status
	}

	buffer := make(map[Body][]Point, len(s.bodies))
	bodies := make([]Body, len(s.bodies))
	length := 0
	for _, b := range s.bodies {
		buffer[b] = make([]Point, 0)
		bodies[length] = b
		length++
	}

	for i := 0; i < length-1; i++ {
		b1 := bodies[i]
		for j := i + 1; j < length; j++ {
			b2 := bodies[j]
			diff := b1.Grav(b2)
			buffer[b1] = append(buffer[b1], diff)
			buffer[b2] = append(buffer[b2], diff.Mul(-1))
		}
	}

	for b, incs := range buffer {
		for _, inc := range incs {
			b.SetInertia(b.GetInertia().Add(inc))
		}
	}

	for b := range buffer {
		if err := b.Move(dt); err != nil {
			s.status = err
			return err
		}
	}

	return nil
}

func (s system) TotalMass() float64 {
	var mass float64
	for _, b := range s.bodies {
		mass += b.GetMass()
	}
	return mass
}

func (s system) String() string {
	c := len(s.bodies)

	pl := ""
	if c > 1 {
		pl = "s"
	}

	return fmt.Sprintf(
		"system(%v object%v, total mass of %vKg)",
		c, pl, s.TotalMass(),
	)
}