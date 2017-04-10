#Design notes:
In the past my iterations tried to do things like list, filter, etc. Now the API will be as stupid as possible and be driven by tests.

#API layout:
## Spells
1. Add a spell via posting to `api/spell/add`
2. Delete a spell via posting the uid to `api/spell/delete`


## Characters
1. Add a character via postigin to `api/character/add`
2. Delete a character via posting to `api/character/delete`

# Go Layout
## Storage 
This iteraition will write everything to JSON on file. The files will be stored as their monotonic id in a directory of the struct type they are. 

### Monotonic IDs
Structs will each have an ID property for identification. A separate file named "meta" will store the highest available integer.