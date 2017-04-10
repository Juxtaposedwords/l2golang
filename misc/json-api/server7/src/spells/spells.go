package spells

import (
	"encoding/json"
	"fmt"
	"myData"
	"myHTTP"
	"myThings/Spell"
	"net/http"
	"regexp"

	"strconv"
)

type handler func(*http.Request) ([]byte, error)

const (
	listPattern           = `/api/spells?/$`
	addSpellPattern       = `^/api/spells/add$`
	listSpellLevelPattern = `^/api/spells/[\d]+$`
	maxPostSize           = 24309
	baseURL               = "/api/spells"
)

var dispatch = map[string]handler{
	listSpellPattern:      ListSpells,
	addSpellPattern:       AddSpell,
	listSpellLevelPattern: ListLevelSpells,
}

func Dispatcher(r *http.Request) ([]byte, error) {
	p := r.URL.Path
	for k, v := range dispatch {
		ok, err := regexp.MatchString(k, p)
		if err != nil {
			return nil, err
		}
		if ok {
			return v(r)
		}
	}
	return nil, myHTTP.NotFoundErr
}

func LoadSpells() ([]Spell, error) {
	return myData.ListSpells()
}

func ListSpells(r *http.Request) ([]byte, error) {
	o, err := myData.ListSpells()
	if err != nil {
		return nil, err
	}
	return json.Marshal(o)
}

func ListLevelSpells(r *http.Request) ([]byte, error) {
	t, err := LoadSpells()
	if err != nil {
		return nil, err
	}

	s, err := strconv.Atoi(r.URL.Path[len("/api/spells/"):])

	if err != nil {
		return nil, err
	}
	var o []Spell
	for _, e := range t {
		if e.Level == s {
			o = append(o, e)
		}
	}
	if len(o) != 0 {
		return json.Marshal(o)
	}
	return []byte(fmt.Sprintf("There are no level %d spells...Some say the magic has gone away.", s)), nil
}

func AddSpell(r *http.Request) ([]byte, error) {
	t, err := LoadSpells()
	if err != nil {
		return nil, err
	}
	err = r.ParseMultipartForm(maxPostSize)
	if err != nil {
		return nil, myHTTP.Unprocessable
	}

	l, err := strconv.Atoi(r.FormValue("level"))
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

	if err = myData.PutSpell(n); err != nil {
		return nil, err
	}
	return []byte(""), nil
}
