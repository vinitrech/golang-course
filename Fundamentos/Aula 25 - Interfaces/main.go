package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Println("A área da forma é ", f.area())
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func main() {
	r := retangulo{10, 10}
	c := circulo{15}

	fmt.Println(r.area(), c.area())
}
