package main

import (
	"fmt"

	v1 "github.com/golubkovden/onboarding/builder/internal/api/v1"
	"github.com/golubkovden/onboarding/builder/internal/director"
	"github.com/golubkovden/onboarding/builder/internal/hero"
)

func main() {
	var (
		d         director.Director
		character v1.Hero
	)
	d = director.NewDirector(hero.NewWizardBuilder())
	character = d.CreateHero("Gandalf")
	fmt.Printf("Hero created: %+v\n", character)

	d = director.NewDirector(hero.NewWarriorBuilder())
	character = d.CreateHero("Aragorn")
	fmt.Printf("Hero created: %+v\n", character)
}
