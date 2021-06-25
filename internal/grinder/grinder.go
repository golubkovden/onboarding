package grinder

import (
	"github.com/golubkovden/onboarding/pkg/coffee"
)

// Grinder defines methods for grind coffee beans
type Grinder interface {
	// Grind coffee beans into grounded one
	Grind(count int) coffee.GroundCoffee
}

type grinder struct{}

func (g *grinder) Grind(count int) coffee.GroundCoffee {
	return coffee.GroundCoffee{Amount: count}
}

// NewGrinder returns a new Grinder instance
func NewGrinder() Grinder {
	return &grinder{}
}
