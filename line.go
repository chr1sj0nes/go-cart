package cart

import (
	"math"
)

// Line in 2D space
type Line2 struct {
	U Vector2
	P Point2
}

// Create line, y = mx + c
func NewLine2(m, c float64) *Line2 {
	return &Line2{Vector2{1, m}, Point2{0, c}}
}

// Distance between line and point
func (l *Line2) DistPt(p *Point2) float64 {
	v := Vector2{l.U[1], -l.U[0]}       // vector perpendicular to line
	r := l.P.To(p)                      // vector from point on line to p
	return math.Abs(v.Dot(r)) / v.Len() // project r onto v
}

// Distance between two lines (zero unless the lines are parallel)
func (l1 *Line2) DistLn(l2 *Line2) float64 {
	if l1.U != l2.U {
		return 0
	}

	return l1.DistPt(&l2.P)
}

// Line in 3D space
type Line3 struct {
	U Vector3
	P Point3
}

// Distance between line and point
func (l *Line3) DistPt(p *Point3) float64 {
	p2 := l.P.Add(&l.U) // 2nd point on line
	return p.To(&l.P).X(p.To(p2)).Len() / l.U.Len()
}

// Distance between two lines
func (l1 *Line3) DistLn(l2 *Line3) float64 {
	v := l1.U.X(&l2.U)
	r := l1.P.To(&l2.P)
	return math.Abs(v.Dot(r)) / v.Len()
}
