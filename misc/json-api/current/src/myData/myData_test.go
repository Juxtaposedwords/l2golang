package myData

import (
	"io/ioutil"
	"myThings"
	"os"
	"path/filepath"
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
func resourceDirDeleter() error {
	_ = os.Remove(fs)
}

func TestPutCharacter(t *testing.T) {

}
func TestGetCharacter(t *testing.T) {

}
func TestPutSpell(t *testing.T) {
	t1 := &myThings.Spell{ID: 1, Level: 3, Name: "Lizard", Description: "Turn people into wizards"}
	if err := PutThing1(t1); err != nil {
		t.Errorf("PutThing1(t1) returned %v", err)
		return
	}
}
func TestGetSpell(t *testing.T) {

}
