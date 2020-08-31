# Learning Go

## Played with:
- Http requests (GET)
- CSV Files (Read, Write)
- Types (Structs, Arrays, Slice, String, ints)
- JSON (Marshal)
- Database (Select Statements)
- Functions (Yes Functions are also types)
- Yaml

## Up Next:
- gRPC / Proto Buffers

## Resources I like:
### DB
- http://go-database-sql.org/
- https://bobby-tables.com/go

### YAML
- https://pkg.go.dev/gopkg.in/yaml.v2?tab=doc

### YouTube
- https://www.youtube.com/watch?v=gi7t6Pl9rxE Workshop: Practical Go  (3 hours 50 minutes long)
- https://www.youtube.com/watch?v=iYrMAVSOkkA Building a Config Service in Go (YAML, Go, Microservice)
- https://www.youtube.com/watch?v=rWBSMsLG8po How I write HTTP Web Services after Eight Years


## Notes

"A value can be of more than one type" repeat 10x

A Writer is an interface for all types that have a Write method
A Reader is an interface for all types that have a Read method

A interface is a type. A value can have more than one type.

ioutil.Open opens the actual file and returns a slice of bytes, os.Read returns an address.