package tests

import (
	"testing"

	"math"

	gravity "bitbucket.org/cacilhas/gravity/system"
)

func TestPoint_Get(t *testing.T) {
	point := gravity.NewPoint(1, 2, 3)

	tests := []struct {
		name          string
		expected, got float64
	}{
		{"GetX", 1, point.GetX()},
		{"GetY", 2, point.GetY()},
		{"GetZ", 3, point.GetZ()},
	}

	for _, test := range tests {
		if test.got != test.expected {
			t.Fatalf(
				"[Point.%v] expected %v, got %v",
				test.name, test.expected, test.got,
			)
		}
	}
}

func TestPoint_Hypot(t *testing.T) {
	point1 := gravity.NewPoint(3, 4, 0)
	point2 := gravity.NewPoint(0, 0, 0)
	point3 := gravity.NewPoint(1, 1, 1)

	tests := []struct {
		name          string
		expected, got float64
	}{
		{"(3, 4)", 5, point1.Hypot()},
		{"(0, 0)", 0, point2.Hypot()},
		{"(1, 1, 1)", math.Sqrt(3), point3.Hypot()},
	}

	for _, test := range tests {
		if test.got != test.expected {
			t.Fatalf(
				"[Point.Hypot %v] expected %v, got %v",
				test.name, test.expected, test.got,
			)
		}
	}
}

func TestPoint_TanXY(t *testing.T) {
	point := gravity.NewPoint(1, 2, 3).TanXY()
	tests := []struct {
		name          string
		expected, got float64
	}{
		{"x", -2, point.GetX()},
		{"y", -1, point.GetY()},
		{"z", 3, point.GetZ()},
	}

	for _, test := range tests {
		if test.got != test.expected {
			t.Fatalf(
				"[Point.TanXY %v] expected %v, got %v",
				test.name, test.expected, test.got,
			)
		}
	}
}

func TestPoint_TanXZ(t *testing.T) {
	point := gravity.NewPoint(1, 2, 3).TanXZ()
	tests := []struct {
		name          string
		expected, got float64
	}{
		{"x", -3, point.GetX()},
		{"y", 2, point.GetY()},
		{"z", -1, point.GetZ()},
	}

	for _, test := range tests {
		if test.got != test.expected {
			t.Fatalf(
				"[Point.TanXY %v] expected %v, got %v",
				test.name, test.expected, test.got,
			)
		}
	}
}

func TestPoint_TanYZ(t *testing.T) {
	point := gravity.NewPoint(1, 2, 3).TanYZ()
	tests := []struct {
		name          string
		expected, got float64
	}{
		{"x", 1, point.GetX()},
		{"y", -3, point.GetY()},
		{"z", -2, point.GetZ()},
	}

	for _, test := range tests {
		if test.got != test.expected {
			t.Fatalf(
				"[Point.TanXY %v] expected %v, got %v",
				test.name, test.expected, test.got,
			)
		}
	}
}

func TestPoint_Add(t *testing.T) {
	point1 := gravity.NewPoint(1, 2, 3)
	point2 := gravity.NewPoint(4, 5, 6)
	pointR := point1.Add(point2)

	tests := []struct {
		name          string
		expected, got float64
	}{
		{"x", 5, pointR.GetX()},
		{"y", 7, pointR.GetY()},
		{"z", 9, pointR.GetZ()},
	}

	for _, test := range tests {
		if test.got != test.expected {
			t.Fatalf(
				"[Point.Add %v] expected %v, got %v",
				test.name, test.expected, test.got,
			)
		}
	}
}

func TestPoint_Add2(t *testing.T) {
	point1 := gravity.NewPoint(1, 2, 3)
	pointR := point1.Add2(4, 5)

	tests := []struct {
		name          string
		expected, got float64
	}{
		{"x", 5, pointR.GetX()},
		{"y", 7, pointR.GetY()},
		{"z", 3, pointR.GetZ()},
	}

	for _, test := range tests {
		if test.got != test.expected {
			t.Fatalf(
				"[Point.Add2 %v] expected %v, got %v",
				test.name, test.expected, test.got,
			)
		}
	}
}

func TestPoint_Add3(t *testing.T) {
	point1 := gravity.NewPoint(1, 2, 3)
	pointR := point1.Add3(4, 5, 6)

	tests := []struct {
		name          string
		expected, got float64
	}{
		{"x", 5, pointR.GetX()},
		{"y", 7, pointR.GetY()},
		{"z", 9, pointR.GetZ()},
	}

	for _, test := range tests {
		if test.got != test.expected {
			t.Fatalf(
				"[Point.Add3 %v] expected %v, got %v",
				test.name, test.expected, test.got,
			)
		}
	}
}

func TestPoint_Diff(t *testing.T) {
	point1 := gravity.NewPoint(6, 5, 4)
	point2 := gravity.NewPoint(1, 2, 3)
	pointR := point1.Diff(point2)

	tests := []struct {
		name          string
		expected, got float64
	}{
		{"x", 5, pointR.GetX()},
		{"y", 3, pointR.GetY()},
		{"z", 1, pointR.GetZ()},
	}

	for _, test := range tests {
		if test.got != test.expected {
			t.Fatalf(
				"[Point.Add %v] expected %v, got %v",
				test.name, test.expected, test.got,
			)
		}
	}
}

func TestPoint_Mul(t *testing.T) {
	point1 := gravity.NewPoint(1, 2, 3)
	pointR := point1.Mul(2)

	tests := []struct {
		name          string
		expected, got float64
	}{
		{"x", 2, pointR.GetX()},
		{"y", 4, pointR.GetY()},
		{"z", 6, pointR.GetZ()},
	}

	for _, test := range tests {
		if test.got != test.expected {
			t.Fatalf(
				"[Point.Add3 %v] expected %v, got %v",
				test.name, test.expected, test.got,
			)
		}
	}
}
