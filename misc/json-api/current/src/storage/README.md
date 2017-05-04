# storage outline
The storage package will abstract away all objects for writing to and from storage. 


## JSON storage
### File storage
Each file will be stored in a 'resource' folder, in a sub-directory which is the name of the struct type. Each file will be named with it's id. 

#### Example:
The following character:
```
{ 
id: 1,
level: 3,
name: "Edgar Codd",
race: "Data-Layer"
}
```
would write to: 
`resources/characters/1.json`

## Non-Exposed Function
### `func read(t interface{}, r io.Reader) error `
unmarshals a JSON-tagged struct from an io.Reader.
### `func write(t interface{}, w io.Writer) error `
marshals a JSON-tagged struct into a byte stream and writes it to an io.Writer.

## Exposed functions
### `func PutCharacter(c storage.Character)error`
### `func GetCharacter(u int)(storage.Character, error)`
### `func PutSpell(s storage.Spell)error`
### `func GetSpell(u int)(storage.Spell, error)`