package main

import (
	"fmt"

	"github.com/torikallcode/module-matematika/shapes"
)

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

	Trapesium := shapes.Trapesium{A: 3, B: 3, C: 3, D: 3, Tinggi: 5}
	fmt.Println(Trapesium.Luas())
	fmt.Println(Trapesium.Keliling())
}
