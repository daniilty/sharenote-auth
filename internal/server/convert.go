package server

import "github.com/daniilty/sharenote-auth/internal/core"

func (l *loginRequest) toService() *core.LoginData {
	return &core.LoginData{
		Email:    l.Email,
		Password: l.Password,
	}
}

func (r *registerRequest) toService() *core.UserInfo {
	return &core.UserInfo{
		Email:    r.Email,
		UserName: r.UserName,
		Name:     r.Name,
		Password: r.Password,
	}
}

func convertCoreUserInfoToResponse(u *core.UserInfo) *userInfoResponse {
	return &userInfoResponse{
		ID:             u.ID,
		Email:          u.Email,
		Name:           u.Name,
		UserName:       u.UserName,
		EmailConfirmed: u.EmailConfirmed,
	}
}
