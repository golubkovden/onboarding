package facade

import (
	"github.com/golubkovden/onboarding/pkg/coffee"
)

type facade struct {
	cm coffeemachine
	n  notifier
}

type coffeemachine interface {
	Validate(name string) error
	Make(name string) *coffee.Drink
}

type notifier interface {
	Notify(drink *coffee.Drink)
	Error(name string, err error)
}

func (f *facade) Make(name string) (*coffee.Drink, error) {
	err := f.cm.Validate(name)
	if err != nil {
		f.n.Error(name, err)
		return nil, err
	}

	c := f.cm.Make(name)
	f.n.Notify(c)

	return c, nil
}

func NewFacade(machine coffeemachine, notifier notifier) *facade {
	return &facade{cm: machine, n: notifier}
}
