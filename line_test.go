package cart

import (
	"testing"
)

func TestLine2DistLn(t *testing.T) {
	tests := []struct {
		l1, l2   Line2
		expected float64
	}{
		{*NewLine2(0, 0), *NewLine2(0, 2), 2},
		{Line2{Vector2{0, 1}, Point2{0, 0}}, Line2{Vector2{0, 1}, Point2{3, 0}}, 3},
		{*NewLine2(1, 2), *NewLine2(3, 4), 0},
	}

	for _, tt := range tests {
		if actual := tt.l1.DistLn(&tt.l2); actual != tt.expected {
			t.Errorf("Line2.DistLn: expected = %f, actual = %f", tt.expected, actual)
		}
	}
}

func TestLine3DistPt(t *testing.T) {
	tests := []struct {
		l        Line3
		p        Point3
		expected float64
	}{
		{Line3{Vector3{1, 0, 0}, Point3{0, 0, 0}}, Point3{2, 0, 0}, 0},
		{Line3{Vector3{1, 0, 0}, Point3{1, 2, 3}}, Point3{1, 2, 4}, 1},
	}

	for _, tt := range tests {
		if actual := tt.l.DistPt(&tt.p); actual != tt.expected {
			t.Errorf("Line3.DistPt: expected = %f, actual = %f", tt.expected, actual)
		}
	}
}

func TestLine3DistLn(t *testing.T) {
	tests := []struct {
		l1, l2   Line3
		expected float64
	}{
		{Line3{Vector3{1, 2, 3}, Point3{0, 7, 0}}, Line3{Vector3{4, 5, 6}, Point3{0, 7, 0}}, 0},
		{Line3{Vector3{1, 0, 0}, Point3{0, 0, 0}}, Line3{Vector3{0, 1, 0}, Point3{0, 0, 9}}, 9},
	}

	for _, tt := range tests {
		if actual := tt.l1.DistLn(&tt.l2); actual != tt.expected {
			t.Errorf("Line3.DistLn: expected = %f, actual = %f", tt.expected, actual)
		}
	}
}
