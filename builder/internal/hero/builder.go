package hero

import v1 "github.com/golubkovden/onboarding/builder/internal/api/v1"

// Builder defines hero builder methods
type Builder interface {
	NewHero(name string)
	BuildRace()
	BuildClass()
	BuildAttributes()
	Hero() (hero v1.Hero)
}
