package rect

func (r *Rect) Perimeter() float32 { // function
	r.P = 2 * (r.L + r.B)
	return r.P
}
