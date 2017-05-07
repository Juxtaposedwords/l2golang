# characters outline
The package for creating and getting our generic fantasy characters.

## Exposed functions
### `func AddCharacter(c types.Character)(error)`
### `func UpdateCharacter(c types.Character)(error)`

### `func httpGetCharacter(*http.Request) ([]byte, error)`
This function will serve the single request for one character identified by a UID.

### `func GetCharacter(u int)(types.Character, error)`
This function will serve only
 
### `func DeleteCharacter(u int)(error)`
