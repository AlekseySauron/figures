package mathpkg

import "math"

func (g Square) area() float64 {

	return g.H * g.W

}

func (g Circle) area() float64 {

	return math.Pi * g.R * g.R

}

func (g Rectangle) area() float64 {

	return g.H * g.W

}

func (g Triangle) area() float64 {

	return 1 / 2 * g.H * g.W

}

// func Area(g Geometry) {
// 	return g.Area()
// }
