package characters

import (
	"encoding/json"
	"fmt"
	"myHTTP"
	"net/http"
	"regexp"
	"storage"
	"strconv"
	"types"
)

const (
	addCharPattern  = `^/api/characters/add$`
	getCharPattern  = `^/api/characters/\d+$`
	listCharPattern = `^/api/characters?/$`
	maxPostSize     = 24309
	URLpath         = "/api/characters/"
)

type characterClient interface {
	GetCharacter(int) (types.Character, error)
	PutCharacter(c types.Character) error
	DeleteCharacter(id int) error
}
type handler func(*http.Request) ([]byte, error)

// map the path to a regex for a function
var dispatch = map[string]handler{
	getCharPattern: GetCharacter,
	addCharPattern: AddCharacter,
}

func init() {
	//	storageClient := storage.NewClient()

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

func HTTPAddCharacter(r *http.Request) ([]byte, error) {

	return nil, nil
}
func AddCharacter(r *http.Request) ([]byte, error) {

	storageClient := storage.NewClient()
	char, err := parseCharacter(r)
	if err != nil {
		return nil, err
	}
	return nil, storageClient.PutCharacter(char)
}

func parseCharacter(r *http.Request) (types.Character, error) {
	if err := r.ParseMultipartForm(maxPostSize); err != nil {
		return types.Character{}, err
	}
	level, err := strconv.Atoi(r.FormValue("level"))
	if err != nil {
		fmt.Printf("%+v\n", r)
		return types.Character{}, err
	}

	name := r.FormValue("name")
	race := r.FormValue("race")
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		return types.Character{}, err
	}
	return types.Character{ID: id, Name: name, Race: race, Level: level}, nil

}
func UpdateCharacter(c types.Character) error {
	return nil
}
func GetCharacter(r *http.Request) ([]byte, error) {
	storageClient := storage.NewClient()
	id, err := strconv.Atoi(r.URL.Path[len(URLpath):])
	if err != nil {
		return nil, err
	}
	char, err := storageClient.GetCharacter(id)
	if err != nil {
		return nil, err
	}
	return json.Marshal(char)
}

func DeleteCharacter(c types.Character) error {
	return nil
}

func buildCharacter(r *http.Request) (types.Character, error) {
	return types.Character{}, nil
}
