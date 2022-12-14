package mathpkg

import "math"

func (g Square) Area() float64 {

	return g.H * g.W

}

func (g Circle) Area() float64 {

	return math.Pi * g.R * g.R

}

func (g Rectangle) Area() float64 {

	return g.H * g.W

}

func (g Triangle) Area() float64 {

	return 0.5 * g.H * g.W

}

func Measure(g Geometry) float64 {
	return g.Area()
}
