package shapes

type LayangLayang struct {
	S1, S2, D1, D2 float64
}

func (l LayangLayang) Luas() float64 {
	return 0.5 * l.D1 * l.D2
}

func (l LayangLayang) Keliling() float64 {
	return 2 * (l.S1 + l.S2)
}
