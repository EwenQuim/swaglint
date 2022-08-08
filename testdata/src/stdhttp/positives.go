package stdhttp

import (
	"net/http"
)

func _(w http.ResponseWriter, r *http.Request) { // want "should have a swagger documentation"
	w.Write([]byte("Hello, world!"))
}

// not documented the right way
func _(w http.ResponseWriter, r *http.Request) { // want "should have the following tags: @Router, @Summary, @Tags"
	w.Write([]byte("Hello, world!"))
}

// @Router /hello [get]
func _(w http.ResponseWriter, r *http.Request) { // want "should have the following tags: @Summary, @Tags"
	w.Write([]byte("Hello, world!"))
}

// @Summary does this thing
// @Tags user
func _(w http.ResponseWriter, r *http.Request) { // want "should have the following tags: @Router"
	w.Write([]byte("Hello, world!"))
}

// @Summary Hello, world!
// @Router /hello [get]
func _(w http.ResponseWriter, r *http.Request) { // want "should have the following tags: @Tags"
	w.Write([]byte("Hello, world!"))
}
