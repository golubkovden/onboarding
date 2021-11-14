package facade

import (
	"github.com/golubkovden/onboarding/facade/internal/api/v1"
)

type catalog interface {
	Get(name string) (v1.Recipe, error)
}

type grinder interface {
	Grind(count int) v1.GroundCoffee
}

type dripper interface {
	Drip(g v1.GroundCoffee, water int) (coffee int)
}

// Coffeemaker defines methods of brewing coffee
type Coffeemaker interface {
	Make(name string) (v1.Drink, error)
}

type coffeemaker struct {
	grinder grinder
	dripper dripper
	catalog catalog
}

func (c *coffeemaker) Make(name string) (drink v1.Drink, err error) {
	r, err := c.catalog.Get(name)
	if err != nil {
		return
	}
	g := c.grinder.Grind(r.CoffeeCount)
	v := c.dripper.Drip(g, r.WatterCount)

	// sets drink attributes
	drink.Name = r.Name
	drink.Volume = v
	return
}

// NewCoffeemaker returns a new Coffeemaker instance
func NewCoffeemaker(catalog catalog, grinder grinder, dripper dripper) Coffeemaker {
	return &coffeemaker{catalog: catalog, grinder: grinder, dripper: dripper}
}
