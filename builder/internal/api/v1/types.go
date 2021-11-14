package v1

// Hero information
type Hero struct {
	Name      string
	Race      string
	Class     string
	Attribute Attribute
}

// Attribute stats
type Attribute struct {
	Intelligence int
	Strength     int
}
