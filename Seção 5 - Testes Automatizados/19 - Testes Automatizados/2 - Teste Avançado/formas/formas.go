package formas

import (
	"math"
)

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	raio float64
}

func (c Circulo) Area() float64 {
	var pi float64 = math.Pi
	var rQuadrado = math.Pow(c.raio, 2)
	var a float64 = pi * rQuadrado
	return a
}
