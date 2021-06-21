package notifier

import (
	"fmt"

	"github.com/golubkovden/onboarding/pkg/coffee"
)

type Service interface {
	Notify(drink *coffee.Drink)
	Error(name string, err error)
}

type logNotifier struct{}

func (l *logNotifier) Notify(drink *coffee.Drink) {
	fmt.Printf("[INFO] %s is ready\n", drink.Name)
}

func (l *logNotifier) Error(name string, err error) {
	fmt.Printf("[ERROR] failed to make %s: %s\n", name, err)
}

func NewLogNotifier() *logNotifier {
	return &logNotifier{}
}
