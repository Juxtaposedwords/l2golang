package remove

import (
	"testing"
)

func TestRemover(t *testing.T){
	tt := []struct{
		desc string
	}{
		{
		desc: "placeholder",
		},
	}
	for _, tc := range tt {
		tc  := tc // Required for parallel tests: https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721
		t.Run(tc.desc, func(t *testing.T){
			t.Parallel()

		})
	}
}