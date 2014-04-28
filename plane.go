package cart

import (
	"math"
)

// Plane in 3D space
type Plane3 struct {
	U Vector3
	P Point3
}

// Distance between point and plane
func (pn *Plane3) DistPt(pt *Point3) float64 {
	r := pn.P.To(pt)
	return math.Abs(pn.U.Dot(r)) / pn.U.Len()
}
