// Package things defines the things that the application will use.  This
// package primarily defines struct types and their JSON serialization rules.
package myThings

// Characters created
type Character struct {
	ID    int    `json: "id"`
	Name  string `json: "name"`
	Race  string `json: "race"`
	Level int    `json: "level"`
}

// Spells available for casting
type Spell struct {
	ID          int    `json: "id"`
	Level       int    `json: "level"`
	Name        string `json: "name"`
	Description string `json: "description"`
}
