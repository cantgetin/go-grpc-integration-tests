package bookrepository

import "context"

type Repository interface {
	GetBook(ctx context.Context, id int32) (*Book, error)
	ListBooks(ctx context.Context) ([]*Book, error)
}
