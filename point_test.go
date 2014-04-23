package cart

import (
	"testing"
)

func TestPoint2Distance(t *testing.T) {
	tests := []struct {
		p1, p2   Point2
		expected float64
	}{
		{Point2{0, 0}, Point2{3, 4}, 5},
		{Point2{4, 2}, Point2{-1, -10}, 13},
	}

	for _, tt := range tests {
		if actual := tt.p1.Distance(&tt.p2); actual != tt.expected {
			t.Errorf("Point2Distance: expected = %f, actual = %f", tt.expected, actual)
		}
	}
}

func TestPoint3Distance(t *testing.T) {
	tests := []struct {
		p1, p2   Point3
		expected float64
	}{
		{Point3{0, 0, 0}, Point3{1, 2, 2}, 3},
		{Point3{4, 2, 1}, Point3{2, 5, -5}, 7},
	}

	for _, tt := range tests {
		if actual := tt.p1.Distance(&tt.p2); actual != tt.expected {
			t.Errorf("Point3Distance: expected = %f, actual = %f", tt.expected, actual)
		}
	}
}
