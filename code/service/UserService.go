package service

import (
	"context"
)

var UserService = userService{}

type userService struct {
	users []*User
}

// AddUser adds a new user to the userService
func (us *userService) AddUser(ctx context.Context, req *User) (*AddUserResponse, error) {
	userID := req.Id
	us.users = append(us.users, &User{
		Id:    userID,
		Name:  req.Name,
		Email: req.Email,
	})

	return &AddUserResponse{UserId: userID}, nil
}

// GetUser retrieves a user from the userService by ID
func (us *userService) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	for _, user := range us.users {
		if user.Id == req.UserId {
			return user, nil
		}
	}
	return nil, nil
}

// ListUsers returns a list of all users in the userService
func (us *userService) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	return &ListUsersResponse{Users: us.users}, nil
}

func (us *userService) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}
