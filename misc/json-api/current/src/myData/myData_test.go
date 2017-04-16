package myData

import (
	"myThings"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

const (
	fakeFS = "/tmp/myData"
	mode   = 744
)

var (
	resources = []string{"spells", "characters", "meta"}
)

func resourceDirMaker() error {
	err := os.Mkdir(fs, mode)
	if err != nil {
		return err
	}
	for _, resource := range resources {
		dir := filepath.Join(fs, resource)
		err = os.Mkdir(dir, mode)
		if err != nil {
			return err
		}
	}
	//	resourceDirDeleter()
	return nil
}
func resourceDirDeleter() {
	_ = os.Remove(fs)
}

func TestCharacter(t *testing.T) {
	fs = fakeFS
	resourceDirMaker()

	t1 := &myThings.Character{
		ID:    1,
		Level: 3,
		Name:  "Edgar Codd",
		Race:  "Data-Layer"}
	if err := PutCharacter(t1); err != nil {
		t.Errorf("PutCharacter(%v) returned error %v\n", t1, err)
	}
	t2 := &myThings.Character{
		ID: t1.GetID()}
	err := GetCharacter(t2)
	if err != nil {
		t.Errorf("GetCharacter(%v) returned %v\n", t2, err)
	}
	if !reflect.DeepEqual(t1, t2) {
		t.Errorf("Written and stored character differ:\n t1: %+v,\n t2:%+v\n", t1, t2)
	}
}
func TestSpell(t *testing.T) {
	resourceDirMaker()
	fs = fakeFS

	t1 := &myThings.Spell{
		ID:          2,
		Level:       4,
		Name:        "Testify",
		Description: "Get realllllly upset over inconsistent unit tests"}
	if err := PutSpell(t1); err != nil {
		t.Errorf("PutCharacter(%v) returned error %v\n", t1, err)
	}
	t2 := &myThings.Spell{
		ID: t1.GetID()}
	err := GetSpell(t2)
	if err != nil {
		t.Errorf("GetCharacter(%v) returned %v\n", t2, err)
	}
	if !reflect.DeepEqual(t1, t2) {
		t.Errorf("Written and stored character differ:\n t1: %+v,\n t2:%+v\n", t1, t2)
	}
	//	resourceDirDeleter()
}
func TestGetSpell(t *testing.T) {

}
