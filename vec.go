package main

import (
	"math"

	"github.com/dmahlow/go-bytearena/models"
)

type Vec struct {
	vector models.Vector2
}

func (v Vec) X() float64 {
	return v.vector[0]
}

func (v Vec) Y() float64 {
	return v.vector[1]
}

func (v *Vec) Set(x, y float64) {
	v.vector = models.NewVector2(x, y)
}

func (v Vec) Distance() float64 {
	return math.Sqrt(v.vector[0]*v.vector[0] + v.vector[1]*v.vector[1])
}

func (v Vec) Angle() float64 {
	d := math.Atan2(v.X(), v.Y()) * 180

	return d
}
