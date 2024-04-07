package structs

import "math"

type Rectangle struct {
  width float64
  height float64
}

type Circle struct {
  radius float64
}

func Perimeter(rectangle Rectangle) float64 {
  return (rectangle.width + rectangle.height) * 2.0
}

func (rectangle Rectangle) Area() float64 {
  return rectangle.width * rectangle.height
}


func (circle Circle) Area() float64 {
  return math.Pi * math.Pow(circle.radius, 2)
}
