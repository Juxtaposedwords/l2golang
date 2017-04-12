package myData

import (
	"io/ioutil"
	"myThings"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

const (
	fs = "../temporary"
)

var (
	resources = []string{"spells", "characters"}
)

func resourceDirMaker() error {
	err := os.Mkdir(fs)
	if err != nil {
		return err
	}
	for _, resource := range resources {
		dir := filepath.Join(fs, resource)
		err = os.Mkdir(dir)
		if err != nil {
			return err
		}
	}
	if err := os.Mkdir(path, mode); err != nil {
		return err
	}
}
func resourceDirDeleter() {
	_ = os.Remove(fs)
}

func TestCharacter(t *testing.T) {
	t1 := &myThings.Character{
		ID:    1,
		Level: 3,
		Name:  "Edgar Codd",
		Race:  "Data-Layer"}
	if err := PutCharacter(t1); err != nil {
		t.Errorf("PutCharacter(%v) returned error %v\n", t1, err)
	}
	t2, err := GetCharacter(t1.id)
	if err != nil {
		t.Errorf("GetCharacter(%v) returned %v\n", t2, err)
	}
	if !reflect.DeepEqual(t1, t2) {
		t.Errorf("Written and stored character differ:\n t1: %v,\n t2:%v\n", t1, t2)
	}
}
func TestSpell(t *testing.T) {
	t1 := &myThings.Spell{
		ID:          2,
		Level:       4,
		Name:        "Testify",
		Description: "Get realllllly upset over inconsistent unit tests"}
	if err := PutSpell(t1); err != nil {
		t.Errorf("PutCharacter(%v) returned error %v\n", t1, err)
	}
	t2, err := GetSpell(t1.id)
	if err != nil {
		t.Errorf("GetCharacter(%v) returned %v\n", t2, err)
	}
	if !reflect.DeepEqual(t1, t2) {
		t.Errorf("Written and stored character differ:\n t1: %v,\n t2:%v\n", t1, t2)
	}
}
func TestGetSpell(t *testing.T) {

}
