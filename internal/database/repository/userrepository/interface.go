package userrepository

import "context"

type Repository interface {
	GetUser(ctx context.Context, id int32) (*User, error)
	ListUsers(ctx context.Context) ([]*User, error)
}
