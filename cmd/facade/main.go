package main

import (
	"fmt"

	v1 "github.com/golubkovden/onboarding/internal/api/v1"
	"github.com/golubkovden/onboarding/internal/catalog"
	"github.com/golubkovden/onboarding/internal/dripper"
	"github.com/golubkovden/onboarding/internal/facade"
	"github.com/golubkovden/onboarding/internal/grinder"
)

func main() {
	recipes := []v1.Recipe{
		{
			Name:        "espresso",
			CoffeeCount: 25,
			WatterCount: 50,
		},
	}

	f := facade.NewCoffeemaker(
		catalog.NewCatalog(recipes),
		grinder.NewGrinder(),
		dripper.NewDripper(),
	)
	drink, err := f.Make("espresso")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Your %s is ready", drink.Name)
}
