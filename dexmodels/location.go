package models

// Location represents a single Location.
type Location struct {
	ID          int    `yaml:"id"`
	Display     string `yaml:"display"`
	CountryCode string `yaml:"country_code"`
	Region      string `yaml:"region"`
}
