package rect

func (r *Rect) Area() float32 { // function
	(*r).A = (*r).L * r.B
	return r.A
}
