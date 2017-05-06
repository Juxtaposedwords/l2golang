package characters

import (
	"errors"
	"testing"
	"types"
)

type fakeCharacterStorageClient struct {
	mapping map[int]error
}

func (cl fakeCharacterStorageClient) GetCharacter(ch *types.Character) error {
	err, ok := cl.mapping[ch.ID]
	if !ok {
		return errEntityNotFound
	}
	return err

}
func (cl fakeCharacterStorageClient) PutCharacter(ch *types.Character) error {
	return nil
}

const (
	existCharID    = 42
	nonExistCharID = 52
)

var (
	errEntityNotFound = errors.New("Entity not found")
)

func TestGetCharacter(t *testing.T) {
	client := fakeCharacterStorageClient{}
	client.mapping = map[int]error{
		42: nil,
		52: errEntityNotFound,
	}
	tt := []struct {
		have *types.Character
		want error
		desc string
	}{
		{
			&types.Character{ID: existCharID},
			nil,
			"Valid get request",
		},
		{
			&types.Character{ID: nonExistCharID},
			errEntityNotFound,
			"Get a non-existent character",
		},
	}
	for _, test := range tt {
		if test.want != client.GetCharacter(test.have) {
			t.Errorf("Failed test case for: have: %+V  want: %+V \n", client.GetCharacter(test.have), test.want)
		}
	}
}
