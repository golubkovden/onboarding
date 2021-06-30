package grinder

import (
	v1 "github.com/golubkovden/onboarding/internal/api/v1"
)

// Grinder defines methods for grind coffee beans
type Grinder interface {
	Grind(count int) v1.GroundCoffee
}

type grinder struct{}

func (g *grinder) Grind(count int) v1.GroundCoffee {
	return v1.GroundCoffee{Amount: count}
}

// NewGrinder returns a new Grinder instance
func NewGrinder() Grinder {
	return &grinder{}
}
