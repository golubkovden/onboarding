package hero

import v1 "github.com/golubkovden/onboarding/builder/internal/api/v1"

type warriorBuilder struct {
	hero v1.Hero
}

func (w *warriorBuilder) NewHero(name string) {
	w.hero = v1.Hero{
		Name: name,
	}
}

func (w *warriorBuilder) BuildRace() {
	w.hero.Race = "Human"
}

func (w *warriorBuilder) BuildClass() {
	w.hero.Class = "Warrior"
}

func (w *warriorBuilder) BuildAttributes() {
	w.hero.Attribute = v1.Attribute{
		Intelligence: 60,
		Strength:     90,
	}
}

func (w *warriorBuilder) Hero() (hero v1.Hero) {
	hero = w.hero
	return
}

// NewWarriorBuilder returns new builder of warriors
func NewWarriorBuilder() Builder {
	return &warriorBuilder{}
}
