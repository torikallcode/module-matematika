package shapes

type Lingkaran struct {
	Jari_jari float64
}

func (l Lingkaran) Luas() float64 {
	pi := 3.14
	return pi * l.Jari_jari * l.Jari_jari
}

func (l Lingkaran) Keliling() float64 {
	pi := 3.14
	return 2 * pi * l.Jari_jari
}

func (l Lingkaran) Diameter() float64 {
	return 2 * l.Jari_jari
}
