package p

import "net/http"

func notDocumented(w http.ResponseWriter, r *http.Request) { // want "should have a swagger documentation"
	w.Write([]byte("Hello, world!"))
}

// Hello, world!
func documented(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
