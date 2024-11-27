package rect

type Rect struct {
	L    float32
	B    float32
	A, P float32
}

func New(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}
