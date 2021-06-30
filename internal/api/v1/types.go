package v1

// Drink is result of coffee maker
type Drink struct {
	Name   string
	Volume int
}

// Recipe is list of ingredients
type Recipe struct {
	Name        string
	CoffeeCount int
	WatterCount int
}

// GroundCoffee is smaller particle of beans
type GroundCoffee struct {
	Amount int
}
