package main

import "testing"

func TestPerimeter(t *testing.T) {
	r := rectangle{10.0, 10.0}
	got := perimeter(r)
	want := 40.0

	if got != want {
		t.Errorf("got '%.2f' want '%.2f'", got, want)
	}
}

func TestArea(t *testing.T) {
	type shape interface {
		area() float64
	}

	areaTests := []struct {
		shape   shape
		hasArea float64
	}{
		{shape: rectangle{12.0, 6.0}, hasArea: 72.0},
		{shape: circle{10}, hasArea: 314.1592653589793},
		{shape: triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.area()
		if got != tt.hasArea {
			t.Errorf("%#v got '%f' want '%f'", tt.shape, got, tt.hasArea)
		}
	}
}
