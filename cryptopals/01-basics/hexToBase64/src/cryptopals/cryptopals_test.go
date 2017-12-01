package cryptopals

import (
	"reflect"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	tt := []struct {
		have []byte
		want []byte
	}{
		{
			[]byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"),
			[]byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"),
		},
		{
			[]byte("4c65742074686520737765657420667265736820627265657a6573206865616c206d652e"),
			[]byte("TGV0IHRoZSBzd2VldCBmcmVzaCBicmVlemVzIGhlYWwgbWUu"),
		},
		{
			[]byte("4c6f6f6b2077696e6461726420616e6420636f6e73696465722050686c65626173"),
			[]byte("TG9vayB3aW5kYXJkIGFuZCBjb25zaWRlciBQaGxlYmFz"),
		},
	}
	for _, tc := range tt {
		if got, err := HexToBase64(tc.have); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Conversion failed. Test: %#v Got: %s err: %s", tc, got, err)
		}

	}
}
func TestXORHex(t *testing.T) {
	tt := []struct {
		have [2][]byte
		want []byte
	}{
		{
			[2][]byte{
				[]byte("1c0111001f010100061a024b53535009181c"),
				[]byte("686974207468652062756c6c277320657965"),
			},
			[]byte("746865206b696420646f6e277420706c6179"),
		},
		{
			[2][]byte{
				[]byte("6d656f77"),
				[]byte("776f6f66"),
			},
			[]byte("1a0a0011"),
		},
	}

	for _, tc := range tt {
		if got, err := XORHex(tc.have[0], tc.have[1]); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Conversion failed.\n TestXORHex(%s, %s)\n Got: %s\n err: %s\n want: %s\n", tc.have[0], tc.have[1], got, err, tc.want)

		}
	}
}
