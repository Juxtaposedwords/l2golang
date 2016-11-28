package main

import "fmt"
import "math"

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
  width, height float64
}

type circle struct {
  radius float64
}

func (r rect) area() float64 {
  return r.width * r.height
}

func (r rect) permi() float64 {
  return 2*r.width + 2*r.height
}

func (c circle) area() float64{
  return math.Pi * c.radius * c.radius
}
