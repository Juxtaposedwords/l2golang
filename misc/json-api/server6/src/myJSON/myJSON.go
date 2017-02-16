package myJSON

import (
	"io/ioutil"
)

func LoadJSON(title string) ([]byte, error) {
	fn := "../resources/" + title + ".json"
	return ioutil.ReadFile(fn)
}

func SaveJSON(title string, b []byte) error {
	fn := "../resources/" + title + ".json"
	return ioutil.WriteFile(fn, b, 0600)

}
