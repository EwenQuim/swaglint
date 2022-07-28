# Swaglint

A linter that helps you to write better swagger documentation.

As there are no code-generated swagger documentation in Go, only comments-based ones, this linter comes in handy and tries to force you to write better documentation.

## Examples

### ❌ BAD:

```go
func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
// REPORTED: swaglint: should have a swagger documentation
```

```go
// WronglyDocumented
// @Summary Hello, world!
// @Tags user
// @Param notName path string true "Name"
// @Router /hello/{name} [get]
func helloWorld(w http.ResponseWriter, r *http.Request) {
    name := chi.URLParam(r, "name")
	w.Write([]byte("Hello, " + name + "!"))
}
// REPORTED: swaglint: param should be named "name" and not "notName"
```

### ✅ GOOD:

```go
// PerfectlyDocumented
// @Summary Hello, world!
// @Tags user
// @Router /hello [get]
func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
```

## Installation & Usage

```bash
go install github.com/EwenQuim/swaglint@latest

swaglint ./...

swaglint package-name

swaglint -h
```

## Roadmap

- [x] Not documented http handlers
- [x] Missing tags
  - [x] Missing summary
  - [x] Missing tags
  - [x] Missing router
- [ ] Type mismatch
  - [ ] Type mismatch in param
  - [ ] Type mismatch in response
- [ ] ⏳ Support for swaggo/swag (parse the comments section with swaggo's internal parser)
- [ ] Support for go-swagger/go-swagger
- [ ] Support for frameworks
  - [ ] ⏳ Support for net/http
  - [ ] ⏳ Support for chi
  - [ ] Support for gin
  - [ ] Support for gorilla/mux
  - [ ] Support for echo
  - [ ] Support for fiber
- [ ] Support for more types of swagger documentation
