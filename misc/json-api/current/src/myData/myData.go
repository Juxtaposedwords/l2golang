package myData

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"myThings"
	"path/filepath"
)

const (
	octalMode = 0664
)

var (
	ErrNotFound      = errors.New("Not found.")
	ErrInvalidType   = errors.New("Invalid struct type.")
	ErrInvalidMode   = errors.New("Invalid access method selected.")
	resourceLocation = "../resources"
	putMap           = map[string]string{
		"*myThings.Character": "characters",
		"*myThings.Spell":     "spells",
	}
	meta      = "meta"
	maxID     = "id.json"
	accessGet = "get"
	accessPut = "put"
)

type object interface {
	GetID() int
	SetID(int)
}

func read(t interface{}, fn string) error {
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &t)
}

func write(t interface{}, fn string) error {
	p, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fn, p, octalMode)
}

func access(t object, mode string) error {
	val, ok := putMap[fmt.Sprintf("%T", t)]
	if !ok {
		return ErrInvalidType
	}
	filename := fmt.Sprintf("%d.json", t.GetID())
	f := filepath.Join(resourceLocation, val, filename)
	switch mode {
	case accessPut:
		if t.GetID() == 0 {
			assignID(t)
		}
		return write(t, f)
	case accessGet:
		return read(t, f)
	default:
		return ErrInvalidMode
	}
	return nil
}

func assignID(t object) error {
	objType := fmt.Sprintf("%T", t)
	mapper := map[string]int{}
	f := filepath.Join(resourceLocation, meta, maxID)
	if err := read(&mapper, f); err != nil {
		return err
	}

	max := mapper[objType]
	mapper[objType] += 1

	if err := write(mapper, f); err != nil {
		return err
	}

	t.SetID(max)

	return nil
}
func PutCharacter(c *myThings.Character) error {
	return access(c, accessPut)
}

func GetCharacter(c *myThings.Character) error {
	return access(c, accessGet)
}
func PutSpell(s *myThings.Spell) error {
	return access(s, accessPut)
}

func GetSpell(s *myThings.Spell) error {
	return access(s, accessGet)
}
