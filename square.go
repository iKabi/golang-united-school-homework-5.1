package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	ia := int(s.a)
	return Point{
		x : s.start.x + ia,
		y : s.start.y + ia,
	}
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return 4 * s.a
}
