package diagClear

import (
	"strings"
	"testing"
)

func TestDiagValue(t *testing.T) {
	table := []struct {
		input string
		want  int
	}{
		{
			input: `3
11 2 4
4 5 6
10 8 -12`,
			want: 15,
		},
		{
			input: `2
1 2
1 4`,
			want: 2,
		},
	}

	for _, test := range table {
		r := strings.NewReader(test.input)
		got, err := diagValue(r)
		if err != nil {
			// test messages should indicate what function failed, and what parameters
			// it was passed, if it's possible, as it is in this case.
			//
			// failing that, the message should tell you which test case in the table it
			// was using - either include the index in the message or include a description
			// field in the test case definition.
			t.Errorf("diagValue(%q) returned %s", test.input, err)
			// this continue is good.
			continue
		}
		if got != test.want {
			t.Errorf("diag.Value(%q)\ngot %d\nwant %d", test.input, got, test.want)
		}
	}
}
