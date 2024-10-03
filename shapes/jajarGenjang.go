package shapes

type JajarGenjang struct {
	SisiA, SisiB, SisiC, SisiD, Alas, Tinggi float64
}

func (j JajarGenjang) Luas() float64 {
	return j.Alas * j.Tinggi
}

func (j JajarGenjang) Keliling() float64 {
	return j.SisiA + j.SisiB + j.SisiC + j.SisiD
}
