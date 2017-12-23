package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexToBase64(input []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(dst, input)
	if err != nil {
		return nil, err
	}

	output := make([]byte, base64.StdEncoding.EncodedLen(len(dst)))
	base64.StdEncoding.Encode(output, dst)
	return output, nil
}

func XORHex(a1 []byte, a2 []byte) ([]byte, error) {
	b1, b2 := make([]byte, hex.DecodedLen(len(a1))), make([]byte, hex.DecodedLen(len(a2)))
	_, err := hex.Decode(b1, a1)
	if err != nil {
		return nil, err
	}
	_, err = hex.Decode(b2, a2)
	if err != nil {
		return nil, err
	}
	if len(a1) != len(a2) {
		return nil, fmt.Errorf("Inputs byte array for XORHex have unequal lengths.")
	}
	b3 := make([]byte, len(b1))
	for i := range b1 {
		b3[i] = b1[i] ^ b2[i]
	}

	return []byte(hex.EncodeToString(b3)), nil
}

func XORHexSingleChar(a1 string, c rune) (string, error) {
	b1, err := hex.DecodeString(a1)
	if err != nil {
		return "", err
	}
	b3 := make([]byte, len(b1))
	for i := range b1 {
		b3[i] = b1[i] ^ byte(c)
	}
	return string(b3), nil
}
func stringFreq(a1 string) map[rune]int {
	o := map[rune]int{}
	for _, e := range a1 {
		o[e]++
	}
	return o
}
