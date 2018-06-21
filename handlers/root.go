package handlers

import (
	"net/http"

	"github.com/kiali/kiali/config"
	"github.com/kiali/kiali/status"
)

// swagger:route GET / getStatus
// ---
// Handler to get the status
//
// responses:
//      200: statusInfo
func Root(w http.ResponseWriter, r *http.Request) {
	getStatus(w, r)
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	RespondWithJSONIndent(w, http.StatusOK, status.Get())
}

func GetToken(w http.ResponseWriter, r *http.Request) {
	u, _, ok := r.BasicAuth()
	if !ok {
		RespondWithJSONIndent(w, http.StatusInternalServerError, u)
		return
	}
	token, error := config.GenerateToken(u)
	if error != nil {
		RespondWithJSONIndent(w, http.StatusInternalServerError, error)
		return
	}
	RespondWithJSONIndent(w, http.StatusOK, token)
}
