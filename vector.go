package cart

import (
	"fmt"
	"math"
)

// Vector in 2D space
type Vector2 [2]float64

// Divide vector by scalar
func (v *Vector2) Div(a float64) *Vector2 {
	return &Vector2{v[0] / a, v[1] / a}
}

// Vector dot product
func (v1 *Vector2) Dot(v2 *Vector2) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1]
}

// Length (Euclidean norm) of vector
func (v *Vector2) Len() float64 {
	return math.Sqrt(v.Dot(v))
}

// Unit (normalized) vector parallel to v (i.e. divide the vector by its length)
func (v *Vector2) Unit() *Vector2 {
	return v.Div(v.Len())
}

func (v *Vector2) String() string {
	return fmt.Sprintf("[%f, %f]", v[0], v[1])
}

// Vector in 3D space
type Vector3 [3]float64

// Divide vector by scalar
func (v *Vector3) Div(a float64) *Vector3 {
	return &Vector3{v[0] / a, v[1] / a, v[2] / a}
}

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

// Length (Euclidean norm) of vector
func (v *Vector3) Len() float64 {
	return math.Sqrt(v.Dot(v))
}

// Unit (normalized) vector parallel to v (i.e. divide the vector by its length)
func (v *Vector3) Unit() *Vector3 {
	return v.Div(v.Len())
}

func (v *Vector3) String() string {
	return fmt.Sprintf("[%f, %f, %f]", v[0], v[1], v[2])
}
