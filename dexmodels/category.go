package models

// Category represents a single Category.
type Category struct {
	ID          int    `yaml:"id"`
	Name        string `yaml:"category_name"`
	Description string `yaml:"description"`
	Priority    int    `yaml:"priority"`
}
