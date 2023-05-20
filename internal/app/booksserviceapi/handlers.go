package booksserviceapi

import (
	"context"
	"go-integration-tests/internal/database/repository/bookrepository"
	"go-integration-tests/pkg/books"
)

func GetBookRepositoryToResponse(
	in *bookrepository.Book) *books.GetBookResponse {
	out := &books.GetBookResponse{}

	out.Book = &books.Book{
		Id:     in.Id,
		Name:   in.Name,
		Author: in.Author,
	}
	return out
}

func ListBooksRepositoryToResponse(
	in []*bookrepository.Book) *books.ListBooksResponse {
	out := &books.ListBooksResponse{}

	out.Books = make([]*books.Book, len(in))
	for i, v := range in {
		out.Books[i] = &books.Book{
			Id:     v.Id,
			Name:   v.Name,
			Author: v.Author,
		}
	}

	return out
}
func (s *ServiceImpl) GetBook(ctx context.Context, req *books.GetBookRequest) (*books.GetBookResponse, error) {
	book, err := s.serviceRepository.GetBook(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return GetBookRepositoryToResponse(book), nil
}

func (s *ServiceImpl) ListBooks(ctx context.Context, req *books.ListBooksRequest) (*books.ListBooksResponse, error) {
	listBooks, err := s.serviceRepository.ListBooks(ctx)
	if err != nil {
		return nil, err
	}
	return ListBooksRepositoryToResponse(listBooks), nil
}
