package structs

import "math"

//Shape is a shape
type Shape interface {
	getArea() float64
}

//Rectangle is a rectangle
type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) getArea() float64 {
	return r.width * r.height
}

//Circle is a circle
type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return c.radius * c.radius * math.Pi
}

//Triangle is a triangle
type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) getArea() float64 {
	return (t.base * t.height) * 0.5
}

//GetShapeArea it gets the shape area
func GetShapeArea(s Shape) float64 {
	area := s.getArea()
	return area
}

//Perimeter calculate the perimeter
func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}
