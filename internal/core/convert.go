package core

import schema "github.com/daniilty/sharenote-grpc-schema"

func convertUserInfoToAddUser(u *UserInfo) *schema.AddUserRequest {
	passwordHash := getMD5Sum(u.Password)

	return &schema.AddUserRequest{
		Email:        u.Email,
		Name:         u.Name,
		UserName:     u.UserName,
		PasswordHash: passwordHash,
	}
}

func convertPBUserToUserInfo(u *schema.User) *UserInfo {
	return &UserInfo{
		ID:             u.Id,
		Email:          u.Email,
		Name:           u.Name,
		UserName:       u.UserName,
		EmailConfirmed: u.EmailConfirmed,
	}
}
