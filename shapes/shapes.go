package shapes

import (
	"fmt"
)

type Shapes interface {
	Luas() float64
	Keliling() float64
}

func CetakInfoShape(shape Shapes) {
	fmt.Printf("Luas: %.2f\n", shape.Luas())
	fmt.Printf("Keliling: %.2f\n", shape.Keliling())
}
