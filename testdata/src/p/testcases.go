package p

import "net/http"

func notDocumented(w http.ResponseWriter, r *http.Request) { // want "should have a swagger documentation"
	w.Write([]byte("Hello, world!"))
}

// Hello, world!
// not documented the right way
func documented(w http.ResponseWriter, r *http.Request) { // want "no @Router tag found"
	w.Write([]byte("Hello, world!"))
}

// Hello, world!
// @Router /hello [get]
func documentedWithSwaggerRouter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
