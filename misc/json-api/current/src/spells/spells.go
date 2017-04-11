package spells

import (
	"encoding/json"
	"myThings"
	"net/http"
)

const (
	listCharPattern      = `^/api/characters?/$`
	addCharPattern       = `^/api/characters/add$`
	listCharLevelPattern = `^/api/characters/\d+$`
	maxPostSize          = 24309
	URLpath              = "/api/characters/"
)

// list all the spells
func SpellList(r *http.Request) ([]byte, error) {
	t := []myThings.Spell{
		{Level: 1, Name: "loud", Description: "Double the decibel, but no higher than 11."},
		{Level: 2, Name: "frustrate", Description: "You speak for hours about the liberal agenda"},
	}
	return json.Marshal(t)
}
