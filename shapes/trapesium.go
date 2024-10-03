package shapes

type Trapesium struct {
	A, B, C, D, Tinggi float64
}

func (t Trapesium) Luas() float64 {
	return ((t.A + t.B) / 2) * t.Tinggi
}

func (t Trapesium) Keliling() float64 {
	return t.A + t.B + t.C + t.D
}
