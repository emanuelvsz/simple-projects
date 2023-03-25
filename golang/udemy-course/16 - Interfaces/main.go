package main

import "fmt"

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Println("A area foi: ", f.area())
}

func (r Retangulo) area() float64 {
	return r.altura * r.largura
}

type Retangulo struct {
	altura  float64
	largura float64
}

type Circulo struct {
	area float64
}

func main() {
	r := Retangulo{10, 20}
	escreverArea(r)
}
