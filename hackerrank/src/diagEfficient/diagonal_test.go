package diagEfficient
import (
	"testing"
	"strings"
	"fmt"
)

func TestDiagValue(t *testing.T){
	table := []struct {
		text  string
		diagValue int
	}{
	{`3
	11 2 4
	4 5 6
	10 8 -12`, 15}}

	fmt.Printf("value: %z ,err: %z", x, err)	

	for _, test := range table {
		r := strings.NewReader(test)
		x,err := diagValue(r)
		size := Version(test.version).PatternSize()
		if size != test.expected {
			t.Errorf("Version %2d, expected %3d but got %3d",
				test.version, test.expected, size)
		}
	}
}
