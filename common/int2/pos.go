package int2

type Int2 struct {
	X, Y int
}

func Add(a Int2, b Int2) Int2 {
	return Int2{a.X + b.X, a.Y + b.Y}
}

type Pos = Int2
type Delta = Int2

func (pos Pos) Moved(delta Delta) Pos {
	return Add(pos, delta)
}
