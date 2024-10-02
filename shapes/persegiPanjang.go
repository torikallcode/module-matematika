package shapes

type PersegiPanjang struct {
	Panjang, Lebar float64
}

func (p PersegiPanjang) Luas() float64 {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() float64 {
	return 2 * (p.Panjang + p.Lebar)
}
