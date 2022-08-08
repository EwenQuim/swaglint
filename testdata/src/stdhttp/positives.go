package stdhttp

import (
	"encoding/json"
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

// @Summary Hello, world!
// @Tags user
// @Router /hello [get]
func _(w http.ResponseWriter, r *http.Request) { // want "clientID is in code but not in docs"
	clientID := r.FormValue("clientID")
	json.NewEncoder(w).Encode(clientID)
}

// @Summary Hello, world!
// @Tags user
// @Param clientID query string true "Client ID"
// @Router /hello [get]
func _(w http.ResponseWriter, r *http.Request) { // want "clientID is in docs but not in code"
	w.Write([]byte("Hello, world!"))
}
