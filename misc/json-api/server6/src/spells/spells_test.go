package spells

import (
	"net/http"
	"reflect"
	"testing"
)

func TestDispatcher(t *testing.T) {
	tests := []struct {
		url  string
		p    string
		h    handler
		want []byte
	}{
		{
			"/api/spells",
			`^/api/spells/add$`,
			func(*http.Request) ([]byte, error) { return []byte("one"), nil },
			[]byte("one"),
		},
		{
			"/api/spells/add",
			`^/api/spells$`,
			func(*http.Request) ([]byte, error) { return []byte("two"), nil },
			[]byte("two"),
		},
	}

	for _, tt := range tests {
		r, err := http.NewRequest("GET", tt.url, nil)
		spells.dispatch[tt.p] = tt.h
		got, err := Dispatcher(r)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Dispatcher(%s), got %s, want %s", tt.url, got, tt.want)
		}
	}
}
