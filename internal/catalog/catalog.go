package catalog

import (
	"errors"
	"strings"

	"github.com/golubkovden/onboarding/pkg/coffee"
)

var (
	ErrRecipeNotFound = errors.New("recipe not found")
)

// Catalog defines methods for resolving recipe of drink
type Catalog interface {
	// Get recipe by name
	Get(name string) (coffee.Recipe, error)
}

type catalog struct {
	recipes []coffee.Recipe
}

func (c *catalog) Get(name string) (coffee.Recipe, error) {
	for _, recipe := range c.recipes {
		if strings.EqualFold(recipe.Name, name) {
			return recipe, nil
		}
	}

	return coffee.Recipe{}, ErrRecipeNotFound
}

// NewCatalog returns a new Catalog instance
func NewCatalog(recipes []coffee.Recipe) Catalog {
	return &catalog{recipes: recipes}
}
