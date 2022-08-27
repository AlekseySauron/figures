package mathpkg

type Square struct {
	H, W float64
}

type Circle struct {
	R float64
}

type Rectangle struct {
	H, W float64
}

type Triangle struct {
	H, W float64
}

type Geometry interface {
	//mathpkg.
	Area() float64
}
