package cart

import (
	"fmt"
)

// Point in 2D space
type Point2 [2]float64

// Vector p1 -> p2
func (p1 *Point2) To(p2 *Point2) *Vector2 {
	return &Vector2{p2[0] - p1[0], p2[1] - p1[1]}
}

// Add vector to point
func (p *Point2) Add(v *Vector2) *Point2 {
	return &Point2{p[0] + v[0], p[1] + v[1]}
}

// Distance between points
func (p1 *Point2) DistPt(p2 *Point2) float64 {
	return p1.To(p2).Len()
}

func (p *Point2) String() string {
	return fmt.Sprintf("(%f, %f)", p[0], p[1])
}

// Point in 3D space
type Point3 [3]float64

// Vector p1 -> p2
func (p1 *Point3) To(p2 *Point3) *Vector3 {
	return &Vector3{p2[0] - p1[0], p2[1] - p1[1], p2[2] - p1[2]}
}

// Add vector to point
func (p *Point3) Add(v *Vector3) *Point3 {
	return &Point3{p[0] + v[0], p[1] + v[1], p[2] + v[2]}
}

// Distance between points
func (p1 *Point3) DistPt(p2 *Point3) float64 {
	return p1.To(p2).Len()
}

func (p *Point3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", p[0], p[1], p[2])
}
