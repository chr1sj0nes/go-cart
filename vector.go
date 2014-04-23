package cart

import ()

// Vector in 2D space
type Vector2 [2]float64

// Vector dot product
func (v1 *Vector2) Dot(v2 *Vector2) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1]
}

// Vector in 3D space
type Vector3 [3]float64

// Vector dot product
func (v1 *Vector3) Dot(v2 *Vector3) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

// Vector cross product
func (v1 *Vector3) X(v2 *Vector3) *Vector3 {
	vx := v1[1]*v2[2] - v1[2]*v2[1]
	vy := v1[0]*v2[2] - v1[2]*v2[0]
	vz := v1[0]*v2[1] - v1[1]*v2[0]
	return &Vector3{vx, vy, vz}
}
