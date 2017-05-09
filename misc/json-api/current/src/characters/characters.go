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
	addCharPattern    = `^/api/characters/add$`
	updateCharPattern = `^/api/characters/update$`
	deleteCharPattern = `^/api/characters/delete/\d+$`
	getCharPattern    = `^/api/characters/\d+$`
	listCharPattern   = `^/api/characters?/$`
	maxPostSize       = 24309
	URLpath           = "/api/characters/"
)

type characterClient interface {
	GetCharacter(int) (types.Character, error)
	PutCharacter(c types.Character) error
	DeleteCharacter(id int) error
}
type handler func(*http.Request) ([]byte, error)

// map the path to a regex for a function
var dispatch = map[string]handler{

	getCharPattern:    GetCharacter,
	addCharPattern:    AddCharacter,
	updateCharPattern: UpdateCharacter,
	deleteCharPattern: DeleteCharacter,
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

func AddCharacter(r *http.Request) ([]byte, error) {
	if r.Method != http.MethodPost {
		return nil, myHTTP.WrongMethod
	}
	storageClient := storage.NewClient()
	char, err := parseCharacter(r)
	if err != nil {
		return nil, err
	}
	if char.ID != 0 {
		return nil, myHTTP.Unprocessable
	}
	if err = storageClient.PutCharacter(char); err != nil {
		return nil, err
	}
	return nil, myHTTP.Created
}

func parseCharacter(r *http.Request) (types.Character, error) {
	decoder := json.NewDecoder(r.Body)
	var char types.Character
	err := decoder.Decode(&char)
	if err != nil {
		return types.Character{}, myHTTP.Unprocessable
	}
	return char, nil
}
func UpdateCharacter(r *http.Request) ([]byte, error) {
	if r.Method != http.MethodPut {
		return nil, myHTTP.WrongMethod
	}
	storageClient := storage.NewClient()
	char, err := parseCharacter(r)
	if err != nil {
		return nil, err
	}
	if char.ID == 0 {
		return nil, myHTTP.Unprocessable
	}
	if err = storageClient.PutCharacter(char); err != nil {
		return nil, err
	}
	return nil, myHTTP.Accepted
}
func GetCharacter(r *http.Request) ([]byte, error) {
	if r.Method != http.MethodGet {
		return nil, myHTTP.WrongMethod
	}
	storageClient := storage.NewClient()
	id, err := strconv.Atoi(r.URL.Path[len(URLpath):])
	if err != nil {
		return nil, myHTTP.NotAnIntForID
	}
	char, err := storageClient.GetCharacter(id)
	if err != nil {
		return nil, myHTTP.InternalError
	}
	return json.Marshal(char)
}

func DeleteCharacter(r *http.Request) ([]byte, error) {
	if r.Method != http.MethodDelete {
		return nil, myHTTP.WrongMethod
	}
	storageClient := storage.NewClient()
	id, err := strconv.Atoi(r.URL.Path[len(URLpath+"delete/"):])
	fmt.Println(r.URL.Path[len(URLpath+"delete/"):])
	if err != nil {
		return nil, myHTTP.NotAnIntForID
	}
	err = storageClient.DeleteCharacter(id)
	if err != nil {
		return nil, myHTTP.Unprocessable
	}
	return nil, myHTTP.OK
}

func buildCharacter(r *http.Request) (types.Character, error) {
	return types.Character{}, nil
}
