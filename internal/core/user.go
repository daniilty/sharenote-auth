package core

import (
	"context"
	"fmt"

	schema "github.com/daniilty/sharenote-grpc-schema"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserInfo struct {
	ID             string
	Name           string
	UserName       string
	Email          string
	EmailConfirmed bool
	// Used for registration
	Password string
}

func (s *ServiceImpl) GetUserInfo(ctx context.Context, uid string) (*UserInfo, bool, error) {
	resp, err := s.usersClient.GetUser(ctx, &schema.GetUserRequest{
		Id: uid,
	})
	if err != nil {
		if status.Code(err) == codes.InvalidArgument {
			return nil, true, fmt.Errorf("invalid user id")
		}

		return nil, false, err
	}

	user := resp.GetUser()

	return convertPBUserToUserInfo(user), true, nil
}
