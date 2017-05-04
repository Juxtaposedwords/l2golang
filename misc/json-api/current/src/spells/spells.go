package spells

import (
	"encoding/json"
	"things"
	"net/http"
)

const (
	addSpellPattern = `^/api/spells/add$`
	getSpellPattern = `^/api/spells/[\d]+$`

	maxPostSize = 24309
	baseURL     = "/api/spells"
)

var dispatch = map[string]handler{
	addSpellPattern: AddSpell,
	getSpellPattern: GetSpell,
}

func AddSpell(r *http.Request) ([]byte, error) {
	t, err := LoadSpell()
	if err != nil {
		return nil, err
	}
	err = r.ParseMultipartForm(maxPostSize)
	if err != nil {
		return nil, myHTTP.Unprocessable
	}
	id, err := strconv.Atoi(r.URL.Path[len("/api/spells/"):])
	if err != nil {
		return nil, myHTTP.Unprocessable
	}
	n := Spell{
		ID:          0,
		Level:       l,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}
	if n.Name == "" || n.Description == "" {
		return nil, myHTTP.Unprocessable
	}

	if err = storage.PutSpell(n); err != nil {
		return nil, err
	}
	return []byte(""), nil
}
func GetSpell(r *http.Request) ([]byte, error) {

}

func LoadSpell() ([]Spell, error) {
	id, err := strconv.Atoi(r.URL.Path[len("/api/spells/"):])
	if err != nil {
		return nil, error
	}
	return storage.GetSpell(id)
}
