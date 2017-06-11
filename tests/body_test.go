package tests

import (
	"testing"

	gravity "bitbucket.org/cacilhas/gravity/system"
)

func TestGravityBody(t *testing.T) {
	t.Run("null mass", func(t *testing.T) {
		body, err := gravity.NewBody("Null", 0, 3, 4, 5)

		if body != nil {
			t.Fatalf("[NewBody] expected not body, got %v", body)
		}

		if err == nil {

			t.Fatal("[NewBody] error not raised")
		}
	})

	t.Run("negative mass", func(t *testing.T) {
		body, err := gravity.NewBody("Negative", -1, 3, 4, 5)

		if body != nil {
			t.Fatalf("[NewBody] expected not body, got %v", body)
		}

		if err == nil {

			t.Fatal("[NewBody] error not raised")
		}
	})

	t.Run("#Get*", func(t *testing.T) {
		body, _ := gravity.NewBody("Sample", 1, 3, 4, 0)

		if got := body.GetName(); got != "Sample" {
			t.Fatalf("[Body.GetName] expected Sample, got %v", got)
		}

		tests := []struct {
			name          string
			expected, got float64
		}{
			{"GetMass", 1, body.GetMass()},
			{"GetPosition().GetX", 3, body.GetPosition().GetX()},
			{"GetPosition().GetY", 4, body.GetPosition().GetY()},
			{"GetPosition().GetZ", 0, body.GetPosition().GetZ()},
			{"GetInertia().GetX", 0, body.GetInertia().GetX()},
			{"GetInertia().GetY", 0, body.GetInertia().GetY()},
			{"GetInertia().GetZ", 0, body.GetInertia().GetZ()},
		}

		for _, test := range tests {
			if test.got != test.expected {
				t.Fatalf(
					"[Body.%v] expected %v, got %v",
					test.name, test.expected, test.got,
				)
			}
		}
	})

	t.Run("#SetPosition", func(t *testing.T) {
		body, _ := gravity.NewBody("Sample", 1, 3, 4, 0)
		body.SetPosition(body.GetPosition().Mul(2))

		tests := []struct {
			name          string
			expected, got float64
		}{
			{"SetPosition() X", 6, body.GetPosition().GetX()},
			{"SetPosition() Y", 8, body.GetPosition().GetY()},
			{"SetPosition() Z", 0, body.GetPosition().GetZ()},
		}

		for _, test := range tests {
			if test.got != test.expected {
				t.Fatalf(
					"[Body.%v] expected %v, got %v",
					test.name, test.expected, test.got,
				)
			}
		}
	})
}
