package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T){
	tests := [] struct{
		before []tuple
		input tuple
		want []tuple
	}{
		{
			before: []tuple{{1,2}, {3,4}, {5,6}},
			input: tuple{1,6},
			want: []tuple{{1,6},},
		},
		{
			before: []tuple{{4,5}},
			input: tuple{1,6},
			want: []tuple{{1,6}},
		},
		{
			before: []tuple{{4,6}},
			input: tuple{2,5},
			want: []tuple{{2,6}},
		},
	}
	for _, tt := range(tests) {
		got := theFunc(tt.before, tt.input)
		if !reflect.DeepEqual(got, tt.want){
			t.Errorf("theFunc(%v, %v): got %v, want %v\n", tt.before, tt.input, got, tt.want)
		}
	}

}