package facade

import (
	"github.com/golubkovden/onboarding/pkg/coffee"
)

type catalog interface {
	Get(name string) (coffee.Recipe, error)
}

type grinder interface {
	Grind(count int) coffee.GroundCoffee
}

type dripper interface {
	Drip(g coffee.GroundCoffee, water int) (coffee int)
}

// Coffeemaker defines methods of brewing coffee
type Coffeemaker interface {
	// Make creates cup of coffee
	Make(name string) (coffee.Drink, error)
}

type coffeemaker struct {
	grinder grinder
	dripper dripper
	catalog catalog
}

func (c *coffeemaker) Make(name string) (drink coffee.Drink, err error) {
	r, err := c.catalog.Get(name)
	if err != nil {
		return
	}

	g := c.grinder.Grind(r.CoffeeCount)
	v := c.dripper.Drip(g, r.WatterCount)

	drink.Name = r.Name
	drink.Volume = v

	return
}

// NewCoffeemaker returns a new Coffeemaker instance
func NewCoffeemaker(catalog catalog, grinder grinder, dripper dripper) Coffeemaker {
	return &coffeemaker{catalog: catalog, grinder: grinder, dripper: dripper}
}
