package characters

import (
	"encoding/json"
	"myThings"
	"net/http"
)

func CharList(r *http.Request) ([]byte, error) {
	t := []myThings.Character{
		{Level: 1, Name: "Maloy", Race: "Dwarf"},
		{Level: 10, Name: "Claw", Race: "Mountain Lion"},
		{Level: 19, Name: "Clem", Race: "Elf"},
	}
	return json.Marshal(t)
}
