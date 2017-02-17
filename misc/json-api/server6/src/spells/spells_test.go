package spells

import (
	"net/http"
	"reflect"
	"testing"
)

func TestDispatcher(t *testing.T) {
	tests := []struct {
		url  string
		h    handler
		want []byte
	}{
		{
			"/api/spells",
			func(*http.Request) ([]byte, error) { return []byte("one"), nil },
			[]byte("one"),
		},
		{
			"/api/spells/add",
			func(*http.Request) ([]byte, error) { return []byte("two"), nil },
			[]byte("two"),
		},
	}

	for _, tt := range tests {
		r := &http.Request{}
		r.URL.Path = tt.url
		dispatch[tt.url] = tt.h
		got, err := Dispatcher(r)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Dispatcher(%s), got %s, want %s", tt.url, got, tt.want)
		}
	}
}
