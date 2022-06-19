package main

import (
	"fmt"
	"math"
)

// Interface serve para "obrigar/atribuir" métodos a uma struct
type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	var pi float64 = math.Pi
	var rQuadrado = math.Pow(c.raio, 2)
	var a float64 = pi * rQuadrado
	return a
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %.2f\n", f.area())
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
