package spells

import (
	"net/http"
	"testing"
)

func TestDispatcher(t *testing.T) {
	tests := []struct {
		url  string
		p    string
		h    handler
		want string
	}{
		{
			"/api/spells",
			listSpellPattern,
			func(*http.Request) ([]byte, error) { return []byte("one"), nil },
			"one",
		},
		{
			"/api/spells/add",
			addSpellPattern,
			func(*http.Request) ([]byte, error) { return []byte("two"), nil },
			"two",
		},
		{
			"/api/spells/2",
			listSpellLevelPattern,
			func(*http.Request) ([]byte, error) { return []byte("three"), nil },
			"three",
		},
	}

	for _, tt := range tests {
		r, err := http.NewRequest("GET", tt.url, nil)
		dispatch[tt.p] = tt.h
		got, err := Dispatcher(r)
		if err != nil {
			t.Error(err)
		}

		if string(got) != tt.want {
			t.Errorf("Dispatcher(%s), got %s, want %s", tt.url, got, tt.want)
		}
	}
}
