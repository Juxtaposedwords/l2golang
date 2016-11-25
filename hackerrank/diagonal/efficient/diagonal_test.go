package diag
import (
	"testing"
	"strings"
	"fmt"
)

func TestDiagValue(t *testing.T){
	test := `3
	890 1 4
	2 -12 2
	1 2 199`
	r := strings.NewReader(test)
	x,err := diagValue(r)
	fmt.Printf("value: %z ,err: %z", x, err)	
}
