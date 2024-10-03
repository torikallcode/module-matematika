package shapes

type Segitiga struct {
	Sisi, Alas, Tinggi float64
}

func (s Segitiga) Luas() float64 {
	return 0.5 * s.Alas * s.Tinggi
}

func (s Segitiga) Keliling() float64 {
	return 3 * s.Sisi
}
