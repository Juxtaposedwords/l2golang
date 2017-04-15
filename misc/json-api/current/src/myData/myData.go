package myData

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"myThings"
	"os"
	"path/filepath"
)

var (
	ErrNotFound = errors.New("Not found.")
	InvalidType = errors.New("Invalid struct type.")
	fs          = "../resources"
	putMap      = map[string][]string{
		"*myThings.Character": []string{"characters"},
		"*myThings.Spell":     []string{"spells"},
	}
)

type object interface {
	GetID() int
}

func read(t interface{}, r io.Reader) error {
	b := &bytes.Buffer{}
	if _, err := b.ReadFrom(r); err != nil {
		return err
	}
	return json.Unmarshal(b.Bytes(), &t)
}
func write(t interface{}, w io.Writer) error {
	p, err := json.Marshal(t)
	if err != nil {
		return err
	}
	_, err = w.Write(p)
	return err
}

func put(t object) error {
	val, ok := putMap[fmt.Sprintf("%T", t)]
	if !ok {
		return InvalidType
	}
	file := fmt.Sprintf("%d.json", t.GetID())
	absFilePath := filepath.Join(fs, val[0], file)
	f, err := os.Create(absFilePath)
	if err != nil {
		return err
	}
	defer f.Close()
	return write(t, f)
}

func get(t object) error {
	val, ok := putMap[fmt.Sprintf("%T", t)]
	if !ok {
		return InvalidType
	}
	file := fmt.Sprintf("%d.json", t.GetID())
	absFilePath := filepath.Join(fs, val[0], file)
	f, err := os.Open(absFilePath)
	if err != nil {
		return err
	}
	defer f.Close()
	return read(t, f)
}
func PutCharacter(c *myThings.Character) error {
	return put(c)
}

func GetCharacter(c *myThings.Character) error {
	return get(c)
}
func PutSpell(s *myThings.Spell) error {
	return put(s)
}

func GetSpell(s *myThings.Spell) error {
	return get(s)
}
