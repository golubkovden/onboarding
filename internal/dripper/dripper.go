package dripper

import (
	"fmt"

	v1 "github.com/golubkovden/onboarding/internal/api/v1"
)

// Dripper defines methods for mixing water and ground coffee
type Dripper interface {
	Drip(groundCoffee v1.GroundCoffee, waterCount int) (coffee int)
}

type dripper struct{}

func (d *dripper) Drip(groundCoffee v1.GroundCoffee, waterCount int) (coffee int) {
	// some process with mixing ground coffee and watter
	fmt.Printf("mixing %d amount of ground coffee with %d amount of water\n", groundCoffee.Amount, waterCount)
	return waterCount
}

// NewDripper returns a new Dripper instance
func NewDripper() Dripper {
	return &dripper{}
}
