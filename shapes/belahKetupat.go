package shapes

type BelahKetupat struct {
	Sisi, D1, D2 float64
}

func (b BelahKetupat) Luas() float64 {
	return 0.5 * b.D1 * b.D2
}

func (b BelahKetupat) Keliling() float64 {
	return 4 * b.Sisi
}
