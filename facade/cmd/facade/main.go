package main

import (
	"fmt"

	"github.com/golubkovden/onboarding/facade/internal/api/v1"
	"github.com/golubkovden/onboarding/facade/internal/catalog"
	"github.com/golubkovden/onboarding/facade/internal/dripper"
	"github.com/golubkovden/onboarding/facade/internal/facade"
	"github.com/golubkovden/onboarding/facade/internal/grinder"
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
