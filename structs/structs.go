package structs

import "math"

type Rectangle struct {
  width float64
  height float64
}

type Circle struct {
  radius float64
}

type Shape interface {
  Area() float64
  Perimeter() float64
}

func (rectangle Rectangle) Perimeter() float64 {
  return (rectangle.width + rectangle.height) * 2.0
}

func (circle Circle) Perimeter() float64 {
  return 2 * math.Pi * circle.radius
}

func (rectangle Rectangle) Area() float64 {
  return rectangle.width * rectangle.height
}

func (circle Circle) Area() float64 {
  return math.Pi * math.Pow(circle.radius, 2)
}
