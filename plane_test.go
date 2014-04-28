package cart

import (
	"testing"
)

func TestPlane3DistPt(t *testing.T) {
	tests := []struct {
		pn       Plane3
		pt       Point3
		expected float64
	}{
		{Plane3{Vector3{1, 0, 0}, Point3{0, 0, 0}}, Point3{0, 1, 0}, 0},
		{Plane3{Vector3{1, 0, 0}, Point3{0, 0, 0}}, Point3{0, 0, 1}, 0},
		{Plane3{Vector3{1, 0, 0}, Point3{0, 0, 0}}, Point3{1, 0, 0}, 1},
		{Plane3{Vector3{0, 1, 0}, Point3{0, 0, 0}}, Point3{0, 2, 0}, 2},
		{Plane3{Vector3{0, 0, 1}, Point3{0, 0, 0}}, Point3{0, 0, 3}, 3},
		{Plane3{Vector3{0, 0, 1}, Point3{0, 0, 0}}, Point3{1, 2, 3}, 3},
	}

	for _, tt := range tests {
		if actual := tt.pn.DistPt(&tt.pt); actual != tt.expected {
			t.Errorf("Plane3.DistPt: expected = %f, actual = %f", tt.expected, actual)
		}
	}
}
