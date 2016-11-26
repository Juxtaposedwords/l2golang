package diag
import (
	"testing"
	"strings"
)


func TestDiagValue(t *testing.T){
	table := []struct {
		have string
		want int
	}{
	{`3
	11 2 4
	4 5 6
	10 8 -12`, 15},
	{`2
	1 2
	1 4`, 2}}

	for _, test := range table {
		r := strings.NewReader(test.have)
		got,err := diagValue(r)
		if err != nil {
			t.Errorf("Err thrown: %s", err)
			continue 
		}
		if got != test.want {
			t.Errorf("Got %d, want %d", got, test.want)
		}
	}
}
