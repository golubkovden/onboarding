package director

import v1 "github.com/golubkovden/onboarding/builder/internal/api/v1"

type builder interface {
	NewHero(name string)
	BuildRace()
	BuildClass()
	BuildAttributes()
	Hero() v1.Hero
}

// Director manages creation of heroes
type Director interface {
	CreateCharacter(name string) (hero v1.Hero)
}

type director struct {
	builder builder
}

func (d *director) CreateCharacter(name string) v1.Hero {
	d.builder.NewHero(name)
	d.builder.BuildRace()
	d.builder.BuildClass()
	d.builder.BuildAttributes()
	return d.builder.Hero()
}

// NewDirector creates new instance
func NewDirector(builder builder) Director {
	return &director{builder: builder}
}
