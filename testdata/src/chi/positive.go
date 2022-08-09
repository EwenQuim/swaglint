package chi

import (
	"net/http"
)

// deleteUser
// @Summary  delete an user
// @Produce  application/json
// @Success  202
// @Failure  400
// @Tags     Users
// @Param    userID path string true "User ID"
// @Router   /users/{userID} [delete]
func A(w http.ResponseWriter, r *http.Request) { // want "'userID' path param is in docs but not in code"

	w.Write([]byte("Hello, world!"))
}
