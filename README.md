# Swaglint - a linter for swaggo/swag

> And your documentation will always be up-to-date.

[![Go Report Card](https://goreportcard.com/badge/github.com/EwenQuim/swaglint)](https://goreportcard.com/report/github.com/EwenQuim/swaglint)

As there are no code-generated swagger documentation in Go, only comments-based ones, this linter comes in handy and tries to force you to write better documentation. Using this tool, your code will always match your documentation. It can also be used to fix this documentation automatically.

## Examples

### ✅ GOOD:

```go
// PerfectlyDocumented
// @Summary Hello, world!
// @Tags user
// @Param name path string true "Name"
// @Router /hello/{name} [get]
func helloWorld(w http.ResponseWriter, r *http.Request) {
  name := chi.URLParam(r, "name")

  w.Write([]byte("Hello, " + name + "!"))
}
```

### ❌ REPORTED:

```go
func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
// REPORTED: swaglint: should have a swagger documentation
```

```go
// WronglyDocumented
// @Param notName path string true "Name"
// @Router /hello/{name} [get]
func helloWorld(w http.ResponseWriter, r *http.Request) {
  name := chi.URLParam(r, "name")
  w.Write([]byte("Hello, " + name + "!"))
}
// REPORTED: swaglint: should have the following tags: @Summary, @Tags
// REPORTED: swaglint: param should be named "name" and not "notName"
```

## Installation

```bash
go install github.com/EwenQuim/swaglint@latest
```

## Usage

```bash
swaglint ./...

swaglint package-name

swaglint -h
```

## Roadmap

The linter is **working**.

This roadmap is here to help me (and you, if you contribute!) to improve the linter. Even if it is incomplete, you can use it in your projects.

- [x] Detect not documented http handlers
- [x] Detect missing tags
  - [x] Missing @Summary
  - [x] Missing @Tags
  - [x] Missing @Router
- [x] Mismatch in query param
  - [x] For net/http
- [x] Mismatch in path param
  - [x] For chi
  - [ ] For gorilla/mux
  - [ ] For fiber
  - [ ] For gin
- [ ] Mismatch in response body
- [ ] ⏳ Full support for swaggo/swag (parse the comments section with swaggo's internal parser)
- [ ] Support for go-swagger/go-swagger
- [ ] Automatically generate the documentation with `-fix`
