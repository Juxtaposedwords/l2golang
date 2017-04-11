package myData

import (
	"myThings"
	"testing"
)

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
