package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"things"
)

const (
	octalMode = 0664
	meta      = "meta"
	maxID     = "id.json"
	accessGet = "get"
	accessPut = "put"
)

var (
	ErrNotFound      = errors.New("Not found.")
	ErrInvalidType   = errors.New("Invalid struct type.")
	ErrInvalidMode   = errors.New("Invalid access method selected.")
	resourceLocation = "../resources"
	putMap           = map[reflect.Type]string{
		reflect.TypeOf(&things.Character{}): "characters",
		reflect.TypeOf(&things.Spell{}):     "spells",
	}
)

type Client interface {
	GetCharacter(*things.Character) error
}

type idObj interface {
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

func access(t idObj, mode string) error {
	val, ok := putMap[reflect.TypeOf(t)]
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

func assignID(t idObj) error {
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
func PutCharacter(c *things.Character) error {
	return access(c, accessPut)
}

func GetCharacter(c *things.Character) error {
	return access(c, accessGet)
}
func PutSpell(s *things.Spell) error {
	return access(s, accessPut)
}

func GetSpell(s *things.Spell) error {
	return access(s, accessGet)
}
