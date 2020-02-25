package tests

import (
	"testing"

	gravity "github.com/cacilhas/gravity/system"
)

func TestSystem(t *testing.T) {
	t.Run("#Status", func(t *testing.T) {

		system, err := gravity.NewSystem()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got := system.Status(); got != nil {
			t.Fatalf("expect system to start with clean status, but got %v", got)
		}
	})

	t.Run("#GetBody", func(t *testing.T) {
		body1, _ := gravity.NewBody("Assemble 1", 10, 0, 0, 0)
		body2, _ := gravity.NewBody("Assemble 2", 5, 0, 0, 0)
		system, err := gravity.NewSystem(body1, body2)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got := system.Status(); got != nil {
			t.Fatalf("expect system to start with clean status, but got %v", got)
		}

		tests := []struct{ expected, got gravity.Body }{
			{body1, system.GetBody(body1.GetName())},
			{body2, system.GetBody(body2.GetName())},
		}

		for _, test := range tests {
			if test.expected != test.got {
				t.Fatalf("expected %v, got %v", test.expected, test.got)
			}
		}
	})

	t.Run("bodies with same name", func(t *testing.T) {
		body1, _ := gravity.NewBody("Assemble", 10, 0, 0, 0)
		body2, _ := gravity.NewBody("Assemble", 5, 0, 0, 0)
		system, err := gravity.NewSystem(body1, body2)

		if system != nil {
			t.Fatalf("expected no system, got %v", system)
		}

		if err == nil {
			t.Fatal("error not raised")
		}
	})

	t.Run("#AddBody", func(t *testing.T) {
		system, _ := gravity.NewSystem()
		bodies := system.GetBodies()

		if got := len(bodies); got != 0 {
			t.Fatalf("expect empty system, got count %v", got)
		}

		body1, _ := gravity.NewBody("Assemble 1", 10, 0, 0, 0)
		system.AddBody(body1)

		body2, _ := gravity.NewBody("Assemble 2", 5, 0, 0, 0)
		system.AddBody(body2)

		if got := len(bodies); got != 2 {
			t.Fatalf("expect system with 2 bodies, got count %v", got)
		}

		tests := []struct{ expected, got gravity.Body }{
			{body1, bodies[body1.GetName()]},
			{body2, bodies[body2.GetName()]},
		}

		for _, test := range tests {
			if test.expected != test.got {
				t.Fatalf("expected %v, got %v", test.expected, test.got)
			}
		}
	})

	t.Run("#RemoveBody", func(t *testing.T) {
		body1, _ := gravity.NewBody("Assemble 1", 10, 0, 0, 0)
		body2, _ := gravity.NewBody("Assemble 2", 5, 0, 0, 0)
		system, _ := gravity.NewSystem(body1, body2)

		if got := system.RemoveBody(body1); !got {
			t.Fatalf("expect notified remotion of %v", body1)
		}
		if got := system.GetBody(body1.GetName()); got != nil {
			t.Fatalf("expect %v to be removed", body1)
		}
		if got := system.GetBody(body2.GetName()); got != body2 {
			t.Fatalf("%v should not be removed", body2)
		}
		if got := system.RemoveBody(body1); got {
			t.Fatalf("%v alread removed, but renotified", body1)
		}
	})

	t.Run("#TotalMass", func(t *testing.T) {
		system, _ := gravity.NewSystem()

		if got := system.TotalMass(); got != 0 {
			t.Fatalf("expect empty system to be zero mass, got %vKg", got)
		}

		body1, _ := gravity.NewBody("Assemble 1", 10, 0, 0, 0)
		system.AddBody(body1)

		body2, _ := gravity.NewBody("Assemble 2", 5, 0, 0, 0)
		system.AddBody(body2)

		if got := system.TotalMass(); got != 15 {
			t.Fatalf("expect total mass of 15Kg, got %vKg", got)
		}

		body1, _ = gravity.NewBody("Assemble 1", 20, 0, 0, 0)
		body2, _ = gravity.NewBody("Assemble 2", 30, 0, 0, 0)
		system, _ = gravity.NewSystem(body1, body2)

		if got := system.TotalMass(); got != 50 {
			t.Fatalf("expect total mass of 50Kg, got %vKg", got)
		}
	})

	t.Run("#Step", func(t *testing.T) {

		body1, _ := gravity.NewBody("Sun", 100, 0, 0, 0)
		body2, _ := gravity.NewBody("Body 1", 5, 0, 10, 0)
		body3, _ := gravity.NewBody("Body 2", 2, 10, 0, 0)
		system, _ := gravity.NewSystem(body1, body2, body3)

		t.Run("should not raise", func(t *testing.T) {
			err := system.Step(1)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})

		growing := func(a, b, c float64) bool {
			return (a < b) && (b < c)
		}

		t.Run("should update bodies", func(t *testing.T) {
			tests := []struct {
				name  string
				check bool
			}{
				{
					body1.String(),
					growing(0, body1.GetPosition().GetX(), 1e-10) &&
						growing(0, body1.GetPosition().GetY(), 1e-10) &&
						body1.GetPosition().GetZ() == 0 &&
						growing(0, body1.GetInertia().GetX(), 1e-8) &&
						growing(0, body1.GetInertia().GetY(), 1e-8) &&
						body1.GetInertia().GetZ() == 0,
				},
				{
					body2.String(),
					growing(0, body2.GetPosition().GetX(), 1e-10) &&
						growing(9.95, body2.GetPosition().GetY(), 10) &&
						body2.GetPosition().GetZ() == 0 &&
						growing(0, body2.GetInertia().GetX(), 10e-8) &&
						growing(-3.4e-10, body2.GetInertia().GetY(), -3.3e-10) &&
						body2.GetInertia().GetZ() == 0,
				},

				{
					body3.String(),
					growing(9.95, body3.GetPosition().GetX(), 10) &&
						growing(0, body3.GetPosition().GetY(), 1e-10) &&
						body3.GetPosition().GetZ() == 0 &&
						growing(-1.4e-10, body3.GetInertia().GetX(), -1.3e-10) &&
						growing(3.3e-12, body3.GetInertia().GetY(), 3.4e-12) &&
						body3.GetInertia().GetZ() == 0,
				},
			}

			for _, test := range tests {
				if !test.check {
					t.Fatalf("unexpected state: %v", test.name)
				}
			}
		})
	})
}
