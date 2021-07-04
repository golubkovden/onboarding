package catalog

import (
	"errors"
	"strings"

	v1 "github.com/golubkovden/onboarding/internal/api/v1"
)

var (
	ErrRecipeNotFound = errors.New("recipe not found")
)

// Catalog defines methods for resolving recipe of drink
type Catalog interface {
	Get(name string) (v1.Recipe, error)
}

type catalog struct {
	recipes []v1.Recipe
}

func (c *catalog) Get(name string) (recipe v1.Recipe, err error) {
	for _, r := range c.recipes {
		if strings.EqualFold(r.Name, name) {
			recipe = r
			return
		}
	}
	err = ErrRecipeNotFound
	return
}

// NewCatalog returns a new Catalog instance
func NewCatalog(recipes []v1.Recipe) Catalog {
	return &catalog{recipes: recipes}
}
