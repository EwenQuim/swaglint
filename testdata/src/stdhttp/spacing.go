package stdhttp

import (
	"encoding/json"
	"net/http"
)

// Query Params
//
//	@Summary	Hello, world!
//	@Tags		user
//	@Param		clientID	query	string	true	"Client ID"
//	@Router		/hello [get]
func _(w http.ResponseWriter, r *http.Request) {
	clientID := r.FormValue("clientID")
	json.NewEncoder(w).Encode(clientID)
}
