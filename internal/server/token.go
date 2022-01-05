package server

import "net/http"

type accessTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

func (a *accessTokenResponse) writeJSON(w http.ResponseWriter) error {
	return writeJSONResponse(w, http.StatusOK, &a)
}
