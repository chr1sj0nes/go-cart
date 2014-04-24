package cart

import (
    "testing"
)

func TestVectorXProduct(t *testing.T) {
	tests := []struct {
		v1, v2, expected Vector3
	}{
		{Vector3{1, 0, 0}, Vector3{0, 1, 0}, Vector3{0, 0, 1}},
		{Vector3{0, 1, 0}, Vector3{0, 0, 1}, Vector3{1, 0, 0}},
		{Vector3{0, 1, 0}, Vector3{1, 0, 0}, Vector3{0, 0, -1}},
		{Vector3{2, 0, 0}, Vector3{0, 3, 0}, Vector3{0, 0, 6}},
	}

	for _, tt := range tests {
		if actual := tt.v1.X(&tt.v2); *actual != tt.expected {
			t.Errorf("Vector3.X: expected = %s, actual = %s", &tt.expected, actual)
		}
	}
}

