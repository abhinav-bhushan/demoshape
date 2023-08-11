package shape

type Rect struct {
	L, B float32
}

// Method:Method contan a reciever
func (r *Rect) Area() float32 {
	return r.L * r.B
}

// function doesn't contain a reciever
func AreaOf(r *Rect) float32 {
	return r.L * r.B
}
