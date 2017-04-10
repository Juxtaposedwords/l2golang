package characters

import (
	"encoding/json"
	"fmt"
	"log"
	"myHTTP"
	"myJSON"
	"net/http"
	"regexp"
	"strconv"
)

type character struct {
	Name  string `json: "name"`
	Race  string `json: "race"`
	Level int    `json: "level"`
}

type handler func(*http.Request) ([]byte, error)

const (
	listCharPattern      = `^/api/characters?/$`
	addCharPattern       = `^/api/characters/add$`
	listCharLevelPattern = `^/api/characters/\d+$`
	maxPostSize          = 24309
	URLpath              = "/api/characters/"
)

var dispatch = map[string]handler{
	listCharPattern:      listChars,
	addCharPattern:       AddCharacter,
	listCharLevelPattern: listLevelChars,
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
	log.Printf("Broke here %s", r.URL.Path)

	return nil, myHTTP.NotFoundErr
}

func loadChars() ([]character, error) {
	b, err := myJSON.LoadJSON("characters")
	if err != nil {
		return nil, err
	}
	var s []character
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s, nil
}

func saveChars(c []character) error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return myJSON.SaveJSON("characters", b)
}

func listChars(r *http.Request) ([]byte, error) {
	return myJSON.LoadJSON("characters")
}

func getChar(r *http.Request) ([]byte, error) {
	x, err := loadChars()
	if err != nil {
		return nil, err
	}
	char := r.URL.Path[len(URLpath):]
	for _, e := range x {
		if e.Name == char {
			return json.Marshal(e)
		}
	}
	return []byte(fmt.Sprintf("The character %s has not been created yet.", char)), nil

}
func listLevelChars(r *http.Request) ([]byte, error) {
	t, err := loadChars()
	if err != nil {
		return nil, err
	}
	s, err := strconv.Atoi(r.URL.Path[len(URLpath):])
	if err != nil {
		return nil, err
	}
	var o []character
	for _, e := range t {
		if e.Level == s {
			o = append(o, e)
		}
	}
	if len(o) != 0 {
		return json.Marshal(o)
	}
	return []byte(fmt.Sprintf("There are no characters at level %d ...yet", s)), nil

}
func AddCharacter(r *http.Request) ([]byte, error) {
	t, err := loadChars()
	if err != nil {
		log.Printf("testing1")
		return nil, err
	}
	if err = r.ParseMultipartForm(maxPostSize); err != nil {
		log.Printf("testing2")
		return nil, err
	}
	l, err := strconv.Atoi(r.FormValue("level"))
	if err != nil {
		log.Printf("testing3")
		return nil, myHTTP.Unprocessable
	}

	n := character{
		Level: l,
		Name:  r.FormValue("name"),
		Race:  r.FormValue("race"),
	}
	if n.Name == "" || n.Race == "" {
		log.Printf("testing4: %s", n)
		return nil, myHTTP.Unprocessable
	}

	t = append(t, n)
	if err = saveChars(t); err != nil {
		return nil, myHTTP.Unprocessable
	}

	return []byte(""), nil
}
