package stdhttp

import "net/http"

// Perfection
// @Summary Hello, world!
// @Tags user
// @Router /hello [get]
func _(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// returns something, not a standard controller
func _(w http.ResponseWriter, r *http.Request) (user string) {
	return "user"
}
