# myData outline
The myData package will abstract away all objects for writing to and from storage. 


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
race: "Datalayer"
}
```
would write to: 
`resources/characters/1.json`

## Exposed functions
### `func PutCharacter(c Character)(error)`
### `func GetCharacter(u integer)(error)`
### `func PutSpell(s Spell)(error)`
### `func GetSpell(u integer)(error)`