package characters

import (
	"encoding/json"
	"myThings"
	"net/http"
)

const (
	addCharPattern = `^/api/characters/add$`
	maxPostSize    = 24309
	URLpath        = "/api/characters/"
)

// map the path to a regex for a function
var dispatch = map[string]handler{
	addCharPattern: AddCharacter,
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

func AddCharacter(r *http.Request) ([]byte, error) {
	return nil, nil
}
func CharList(r *http.Request) ([]byte, error) {
	t := []myThings.Character{
		{Level: 1, Name: "Maloy", Race: "Dwarf"},
		{Level: 10, Name: "Claw", Race: "Mountain Lion"},
		{Level: 19, Name: "Clem", Race: "Elf"},
	}
	return json.Marshal(t)
}
