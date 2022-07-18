package p

import "net/http"

// returns something, not a standard controller
func extractUser(w http.ResponseWriter, r *http.Request) (user string) {
	return "user"
}
