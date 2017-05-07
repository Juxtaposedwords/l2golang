package storage

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"types"
)

const (
	mode = 0774
)

var (
	fakeResourceLocation = filepath.Join(os.TempDir(), "/myData")
	resources            = []string{"spells", "characters", "meta"}
)

func resourceDirMaker() error {
	for _, resource := range resources {
		dir := filepath.Join(resourceLocation, resource)
		err := os.MkdirAll(dir, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

func makeResourceDir(t *testing.T) string {
	// get the name of the test that called makeResourceDir.
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f, _ := runtime.CallersFrames(pc).Next()
	n := strings.Split(f.Function, ".")[1]

	resourceLocation = filepath.Join(fakeResourceLocation, n)
	for _, resource := range resources {
		dir := filepath.Join(resourceLocation, resource)
		err := os.MkdirAll(dir, mode)
		if err != nil {
			t.Fatal(err)
		}
	}
	return resourceLocation
}

func removeResourceDir(t *testing.T, d string) {
	if err := os.RemoveAll(d); err != nil {
		t.Fatalf("Error removing %q: %s", d, err)
	}
}

func TestReadWrite(t *testing.T) {
	d := makeResourceDir(t)
	defer removeResourceDir(t, d)

	var t1, t2 map[string]int
	t1 = map[string]int{"a": 1, "b": 2}

	fn := filepath.Join(d, "test.json")
	if err := write(t1, fn); err != nil {
		t.Fatalf("write(%v, %q): %s", t1, fn, err)
	}
	if err := read(&t2, fn); err != nil {
		t.Fatalf("read(%q): %s", fn, err)
	}

	if !reflect.DeepEqual(t1, t2) {
		t.Errorf("t1 != t2: %v, %v", t1, t2)
	}
}

func TestCharacter(t *testing.T) {
	resourceLocation := makeResourceDir(t)
	client := NewClient()
	defer removeResourceDir(t, resourceLocation)

	t1 := types.Character{
		ID:    1,
		Level: 3,
		Name:  "Edgar Codd",
		Race:  "Data-Layer"}
	if err := client.PutCharacter(&t1); err != nil {
		t.Errorf("PutCharacter(%v) returned error %v\n", t1, err)
	}
	t2, err := client.GetCharacter(t1.GetID())
	if err != nil {
		t.Errorf("GetCharacter(%v) returned %v\n", t2, err)
	}
	if !reflect.DeepEqual(t1, t2) {
		t.Errorf("Written and stored character differ:\n t1: %+v,\n t2:%+v\n", t1, t2)
	}

}
func TestSpell(t *testing.T) {
	resourceLocation := makeResourceDir(t)
	defer removeResourceDir(t, resourceLocation)

	client := NewClient()
	err := resourceDirMaker()
	if err != nil {
		t.Errorf("Error creating a directory %s", err)
	}
	t1 := types.Spell{
		ID:          2,
		Level:       4,
		Name:        "Testify",
		Description: "Get realllllly upset over inconsistent unit tests"}
	if err = client.PutSpell(&t1); err != nil {
		t.Errorf("PutCharacter(%v) returned error %v\n", t1, err)
	}
	t2, err := client.GetSpell(t1.GetID())
	if err != nil {
		t.Errorf("GetCharacter(%v) returned %v\n", t2, err)
	}
	if !reflect.DeepEqual(t1, t2) {
		t.Errorf("Written and stored character differ:\n t1: %+v,\n t2:%+v\n", t1, t2)
	}
	if err = os.RemoveAll(resourceLocation); err != nil {
		t.Errorf("There was an error removing: %s", err)
	}
}
func ExampleCmd_Output(t *testing.T) {
	out, err := exec.Command("/bin/cat", "/tmp/myData/TestAssign/meta/id.json").Output()
	if err != nil {
		t.Errorf("catting borked: %+v\n", err)
	}
	fmt.Printf("The contents of id.json are: %v\n", string(out))
}
func TestAssignID(t *testing.T) {
	resourceLocation := makeResourceDir(t)
	defer removeResourceDir(t, resourceLocation)

	metaDataMap := map[string]int{
		"*types.Character": 43,
		"*types.Spell":     56}
	metaFilePath := filepath.Join(resourceLocation, "meta", "id.json")
	err := write(metaDataMap, metaFilePath)
	if err != nil {
		t.Errorf("There was an issue with opening the temp config file: %s", err)
		return
	}
	testSpell := &types.Spell{}
	if err = assignID(testSpell); err != nil {
		t.Errorf("There was a problem the assignment: %s", err)
	}
	if testSpell.ID != 56 {
		t.Errorf("There a problem assigning the correct ID, got the ID: %+v", testSpell)
	}
}
