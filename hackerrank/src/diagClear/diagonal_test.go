package diagClear

import (
	"strings"
	"testing"
)

func TestDiagValue(t *testing.T) {
	table := []struct {
		input string
		want  int
		expError error
	}{
		{
			input: `3
11 2 4
4 5 6
10 8 -12`,
			want: 15, 
			expError: nil,
		},
		{
			input: `2
1 2
1 4`,
			want: 2,
			expError: nil,
		},
	}

	for _, test := range table {
		r := strings.NewReader(test.input)
		got, err := diagValue(r)
		switch {
		case err != nil && test.expError != nil :
			t.Errorf("diagValue(%q) returned %s", test.input, err)
		case got != test.want : 
			t.Errorf("diag.Value(%q)\ngot %d\nwant %d", test.input, got, test.want)
		default :
			continue
		}
	}
}
