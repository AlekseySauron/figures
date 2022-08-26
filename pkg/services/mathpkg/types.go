package mathpkg

type Square struct {
	W, H float64
}

type Circle struct {
	R float64
}

type Rectangle struct {
	W, H float64
}

type Triangle struct {
	W, H float64
}

type Geometry interface {
	Area() float64
}
