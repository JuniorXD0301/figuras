package figuras

import "fmt"

type Figura2D interface {
	area() float64
	perimetro() float64
}

func Medidas(f Figura2D) {
	fmt.Println(f)
	fmt.Println(f.area())
	fmt.Println(f.perimetro())
}
