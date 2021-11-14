package hero

import v1 "github.com/golubkovden/onboarding/builder/internal/api/v1"

type wizardBuilder struct {
	hero v1.Hero
}

func (w *wizardBuilder) NewHero(name string) {
	w.hero = v1.Hero{
		Name: name,
	}
}

func (w *wizardBuilder) BuildRace() {
	w.hero.Race = "Maia"
}

func (w *wizardBuilder) BuildClass() {
	w.hero.Class = "wizard"
}

func (w *wizardBuilder) BuildAttributes() {
	w.hero.Attribute = v1.Attribute{
		Intelligence: 90,
		Strength:     50,
	}
}

func (w *wizardBuilder) Hero() (hero v1.Hero) {
	hero = w.hero
	return
}

// NewWizardBuilder returns new builder of wizards
func NewWizardBuilder() Builder {
	return &wizardBuilder{}
}
