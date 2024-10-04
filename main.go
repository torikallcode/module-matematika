package main

import "github.com/torikallcode/module-matematika/shapes"

func main() {
	// hasil := shapes.Persegi{Sisi: 3}
	// fmt.Printf("Luas: %.2f\n", hasil.Luas())
	// fmt.Printf("Keliling: %.2f\n", hasil.Keliling())

	// pp := shapes.PersegiPanjang{Panjang: 3, Lebar: 4}
	// fmt.Printf("Luas: %.2f\n", pp.Luas())
	// fmt.Printf("Keliling: %.2f\n", pp.Keliling())

	// segitiga := shapes.Segitiga{Sisi: 3, Alas: 7, Tinggi: 11}
	// fmt.Println(segitiga.Luas())
	// fmt.Println(segitiga.Keliling())

	// 	JajarGenjang := shapes.JajarGenjang{SisiA: 3, SisiB: 3, SisiC: 3, SisiD: 3, Alas: 3, Tinggi: 4}
	// 	fmt.Println(JajarGenjang.Luas())
	// 	fmt.Println(JajarGenjang.Keliling())

	// Trapesium := shapes.Trapesium{A: 3, B: 3, C: 3, D: 3, Tinggi: 5}
	// fmt.Println(Trapesium.Luas())
	// fmt.Println(Trapesium.Keliling())

	// Lingkaran := shapes.Lingkaran{Jari_jari: 3}
	// fmt.Printf("Luas : %.2f\n", Lingkaran.Luas())
	// fmt.Printf("Keliling : %.2f\n", Lingkaran.Keliling())
	// fmt.Printf("Diameter : %.2f\n", Lingkaran.Diameter())

	// belah_ketupat := shapes.BelahKetupat{Sisi: 3, D1: 3, D2: 3}
	// shapes.CetakInfoShape(belah_ketupat)

	layang_layang := shapes.LayangLayang{S1: 3, S2: 3, D1: 3, D2: 3}
	shapes.CetakInfoShape(layang_layang)
}
