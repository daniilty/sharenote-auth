package server

import (
	"fmt"
	"net/http"
)

type registerRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (r *registerRequest) validate() error {
	if r.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	if r.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if r.UserName == "" {
		return fmt.Errorf("userName cannot be empty")
	}

	if r.Password == "" {
		return fmt.Errorf("password cannot be empty")
	}

	return nil
}

func (h *HTTP) register(w http.ResponseWriter, r *http.Request) {
	resp := h.getRegisterResponse(r)

	resp.writeJSON(w)
}

func (h *HTTP) getRegisterResponse(r *http.Request) response {
	if r.Body == http.NoBody {
		return getBadRequestWithMsgResponse("no body")
	}

	req := &registerRequest{}

	err := unmarshalReader(r.Body, req)
	if err != nil {
		return getBadRequestWithMsgResponse(err.Error())
	}

	err = req.validate()
	if err != nil {
		return getBadRequestWithMsgResponse(err.Error())
	}

	accessToken, ok, err := h.service.Register(r.Context(), req.toService())
	if err != nil {
		if ok {
			return getBadRequestWithMsgResponse(err.Error())
		}

		h.logger.Errorw("Register user.", "err", err)

		return getInternalServerErrorResponse()
	}

	return &accessTokenResponse{
		AccessToken: accessToken,
	}
}
