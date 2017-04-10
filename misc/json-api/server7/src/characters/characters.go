package characters

import (
	"encoding/json"
	"fmt"
	"myData"
	"myHTTP"
	"myThings"
	"net/http"
	"regexp"
	"strconv"
)

type handler func(*http.Request) ([]byte, error)

const (
	listCharPattern      = `^/api/characters?/$`
	addCharPattern       = `^/api/characters/add$`
	listCharLevelPattern = `^/api/characters/\d+$`
	maxPostSize          = 24309
	URLpath              = "/api/characters/"
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

func loadChars() ([]myThings.Character, error) {
	return myData.ListCharacters()
}

func getChar(r *http.Request) ([]byte, error) {
	x, err := loadChars()
	if err != nil {
		return nil, err
	}
	id := r.URL.Path[len(URLpath):]
	for _, e := range x {
		if e.id == id {
			return json.Marshal(e)
		}
	}
	return []byte(fmt.Sprintf("The character %s has not been created yet.", char)), nil

}
func AddCharacter(r *http.Request) ([]byte, error) {
	t, err := loadChars()
	if err != nil {
		return nil, err
	}
	if err = r.ParseMultipartForm(maxPostSize); err != nil {
		return nil, err
	}
	l, err := strconv.Atoi(r.FormValue("level"))
	if err != nil {
		return nil, myHTTP.Unprocessable
	}

	n := myThings.Character{
		ID:    0,
		Level: l,
		Name:  r.FormValue("name"),
		Race:  r.FormValue("race"),
	}
	if n.Name == "" || n.Race == "" {
		return nil, myHTTP.Unprocessable
	}

	if err = myData.PutCharacter(n); err != nil {
		return nil, myHTTP.Unprocessable
	}

	return []byte(""), nil
}
