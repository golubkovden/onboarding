package coffeemachine

import (
	"fmt"
	"strings"

	"github.com/golubkovden/onboarding/pkg/coffee"
)

type Service interface {
	Validate(name string) error
	Make(name string) *coffee.Drink
}

type coffemachine struct {
	available []string
}

func (c *coffemachine) Validate(name string) error {
	for _, a := range c.available {
		if strings.EqualFold(a, name) {
			return nil
		}
	}

	return fmt.Errorf("unavailable type of drink: %s", name)
}

func (c *coffemachine) Make(name string) *coffee.Drink {
	return &coffee.Drink{Name: name}
}

func New(available []string) *coffemachine {
	return &coffemachine{
		available: available,
	}
}
