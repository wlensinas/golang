package main

import (
	"fmt"
	"reflect"
)

type shape interface {
	getArea() float64
}

type name interface {
	getName() string
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLenght float64
}

func main() {
	t := triangle{}
	s := square{}
	t.base = 10
	t.height = 5
	s.sideLenght = 4
	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	t := reflect.TypeOf(s)
	fmt.Println("The Area of", t.Name(), "is ", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}
