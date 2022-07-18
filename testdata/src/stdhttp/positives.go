package stdhttp

import (
	"net/http"
)

func _(w http.ResponseWriter, r *http.Request) { // want "should have a swagger documentation"
	w.Write([]byte("Hello, world!"))
}

// not documented the right way
func _(w http.ResponseWriter, r *http.Request) { // want "no @Summary tag found"
	w.Write([]byte("Hello, world!"))
}

// @Router /hello [get]
func _(w http.ResponseWriter, r *http.Request) { // want "no @Summary tag found"
	w.Write([]byte("Hello, world!"))
}

// @Summary does this thing
// @Tags user
func _(w http.ResponseWriter, r *http.Request) { // want "no @Router tag found"
	w.Write([]byte("Hello, world!"))
}

// @Summary Hello, world!
// @Router /hello [get]
func _(w http.ResponseWriter, r *http.Request) { // want "no @Tags tag found"
	w.Write([]byte("Hello, world!"))
}
