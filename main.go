package main

import (
	"fmt"

	"github.com/torikallcode/module-matematika/shapes"
)

func main() {
	hasil := shapes.Persegi{Sisi: 3}
	fmt.Printf("Luas: %.2f\n", hasil.Luas())
	fmt.Printf("Keliling: %.2f\n", hasil.Keliling())

	pp := shapes.PersegiPanjang{Panjang: 3, Lebar: 4}
	fmt.Printf("Luas: %.2f\n", pp.Luas())
	fmt.Printf("Keliling: %.2f\n", pp.Keliling())
}
