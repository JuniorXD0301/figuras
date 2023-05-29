package figuras

type Circulo struct {
	Radio float64
}

func (c *Circulo) area() float64 {
	return 3.14 * (c.Radio * c.Radio)
}

func (c *Circulo) perimetro() float64 {
	return 2 * 3.14 * c.Radio
}
