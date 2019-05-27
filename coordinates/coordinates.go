package coordinates

import "math"

// CartesianPair represents (x, y) coordinates on a Cartesian plane
type CartesianPair struct {
	X float64
	Y float64
}

// PolarPair represents a polar coordinate
type PolarPair struct {
	R     float64
	Theta float64
}

// Convert a Cartesian coordinate to a polar coordinate
func (p CartesianPair) Convert() PolarPair {
	r := math.Sqrt(p.X*p.X + p.Y*p.Y)

	adjustment := 0

	if p.X < 0 && p.Y >= 0 {
		// Quadrant II
		adjustment = 180
	} else if p.X < 0 && p.Y < 0 {
		// Quadrant III
		adjustment = 180
	} else if p.X > 0 && p.Y < 0 {
		// Quadrant IV
		adjustment = 360
	}

	theta := math.Atan(p.Y/p.X) + float64(adjustment)
	return PolarPair{R: r, Theta: theta}
}
