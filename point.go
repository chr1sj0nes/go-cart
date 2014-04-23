package cart

import (
	"math"
)

// Point in 2D space
type Point2 [2]float64

// Vector p1 -> p2
func (p1 *Point2) To(p2 *Point2) *Vector2 {
	return &Vector2{p2[0] - p1[0], p2[1] - p1[1]}
}

// Distance between points
func (p1 *Point2) Distance(p2 *Point2) float64 {
	v := p1.To(p2)
	return math.Sqrt(v.Dot(v))
}

// Point in 3D space
type Point3 [3]float64

// Vector p1 -> p2
func (p1 *Point3) To(p2 *Point3) *Vector3 {
	return &Vector3{p2[0] - p1[0], p2[1] - p1[1], p2[2] - p1[2]}
}

// Distance between points
func (p1 *Point3) Distance(p2 *Point3) float64 {
	v := p1.To(p2)
	return math.Sqrt(v.Dot(v))
}
