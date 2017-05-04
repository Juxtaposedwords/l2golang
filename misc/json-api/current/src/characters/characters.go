package characters

import (
	"encoding/json"
	"things"
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

func AddCharacter(c things.Character) error {
	return nil
}
func UpdateCharacter(c things.Character) error {
	return nil
}
func GetCharacter(u int) (things.Character, error) {
	return nil, nil
}
func DeleteCharacter(u int) error {
	return nil
}

func buildCharacter(r *http.Request) (things.Character, error) {

}
