package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"types"
)

const (
	octalMode    = 0664
	meta         = "meta"
	maxID        = "id.json"
	accessGet    = "get"
	accessPut    = "put"
	accessDelete = "delete"
	accessUpdate = "update"
)

var (
	ErrNotFound      = errors.New("Not found.")
	ErrInvalidType   = errors.New("Invalid struct type.")
	ErrInvalidMode   = errors.New("Invalid access method selected.")
	resourceLocation = "../resources"
	putMap           = map[reflect.Type]string{
		reflect.TypeOf(&types.Character{}): "characters",
		reflect.TypeOf(&types.Spell{}):     "spells",
	}
)

type idObj interface {
	GetID() int
	SetID(int)
}

type Client struct {
}

func NewClient() *Client {
	return &Client{}
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
	var filename, f string
	if t.GetID() != 0 {
		filename = fmt.Sprintf("%d.json", t.GetID())
		f = filepath.Join(resourceLocation, val, filename)
		if _, err := os.Stat(f); os.IsNotExist(err) {
			return err
		}
	}
	if t.GetID() == 0 {
		assignID(t)
	}
	filename = fmt.Sprintf("%d.json", t.GetID())
	f = filepath.Join(resourceLocation, val, filename)
	// Make sure the file exists

	switch mode {
	case accessPut:
		return write(t, f)
	case accessGet:
		return read(t, f)
	case accessDelete:
		return os.Remove(f)
	case accessUpdate:
		return write(t, f)
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
func (cl *Client) PutCharacter(c types.Character) error {
	return access(&c, accessPut)
}
func (cl *Client) DeleteCharacter(id int) error {
	return access(&types.Character{ID: id}, accessDelete)
}
func (cl *Client) UpdateCharacter(c types.Character) error {
	return access(&c, accessUpdate)
}

func (cl *Client) GetCharacter(id int) (types.Character, error) {
	char := types.Character{ID: id}
	err := access(&char, accessGet)
	if err != nil {
		return types.Character{}, err
	}
	return char, nil
}
func (cl *Client) PutSpell(s *types.Spell) error {
	return access(s, accessPut)
}

func (cl *Client) GetSpell(id int) (types.Spell, error) {
	spell := types.Spell{ID: id}
	err := access(&spell, accessGet)
	if err != nil {
		return types.Spell{}, err
	}
	return spell, nil
}
