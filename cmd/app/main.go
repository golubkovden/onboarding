package main

import (
	"github.com/golubkovden/onboarding/internal/coffeemachine"
	"github.com/golubkovden/onboarding/internal/facade"
	"github.com/golubkovden/onboarding/internal/notifier"
)

func main() {
	available := []string{"espresso", "cappuccino"}
	if err := run(available); err != nil {
		panic(err)
	}
}

func run(available []string) error {
	cm := coffeemachine.New(available)
	n := notifier.NewLogNotifier()
	machine := facade.NewFacade(cm, n)

	_, err := machine.Make("espresso")
	if err != nil {
		return err
	}

	_, err = machine.Make("cappuccino")
	return err
}
