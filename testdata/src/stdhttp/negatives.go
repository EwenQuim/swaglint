package stdhttp

import (
	"encoding/json"
	"net/http"
)

// Perfection
// @Summary Hello, world!
// @Tags user
// @Router /hello [get]
func _(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// Query Params
// @Summary Hello, world!
// @Tags user
// @Param clientID query string true "Client ID"
// @Router /hello [get]
func _(w http.ResponseWriter, r *http.Request) {
	clientID := r.FormValue("clientID")
	json.NewEncoder(w).Encode(clientID)
}

// returns something, not a standard controller
func _(w http.ResponseWriter, r *http.Request) (user string) {
	return "user"
}
