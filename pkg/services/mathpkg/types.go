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
	Area() float64
}

func NewSquare(H float64, W float64) Geometry {
	return &Square{H: H, W: W}
}

func NewCircle(H float64) Geometry {
	return &Circle{R: H}
}

func NewRectangle(H float64, W float64) Geometry {
	return &Rectangle{H: H, W: W}
}

func NewTriangle(H float64, W float64) Geometry {
	return &Triangle{H: H, W: W}
}
