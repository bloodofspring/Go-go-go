package test_archive

import "fmt"

type Shape interface { // Собственный "паттерн" методов объекта, условные "ожидания" относительно его методов
	area() float64
	perimeter() float64
}

type multiShape struct {
	shapes []Shape
}

type square struct {
	width float64
}

func (s square) area() float64 {
	return s.width * s.width
}

func (s square) perimeter() float64 {
	return s.width * 4
}

type rect struct {
	square // наследование
	height float64
}

func (r rect) area() float64 { // Переопределение метода
	return r.square.width * r.height
}

func (r rect) perimeter() float64 { // Переопределение метода
	return r.square.width*2 + r.height*2
}

func (m multiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func TestWhatAboutInterfaces() {
	someMultiShape := multiShape{[]Shape{
		rect{square{1}, 2},
		rect{square{2}, 1.5},
		square{width: 2},
	}}
	fmt.Println(someMultiShape.area())
}
