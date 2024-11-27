package square

type Square struct { // exported
	S    float32 // exporeted, unrestricted // similar to public
	a, p float32 // restricted, unexporeted // similar to private
}

func (s *Square) Area() float32 {
	s.a = s.S * s.S
	return s.a
}
func (s *Square) Perimeter() float32 {
	s.p = 4 * s.S
	return s.p
}
