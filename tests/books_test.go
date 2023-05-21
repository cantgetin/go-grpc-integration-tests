package tests

import (
	"context"
	"go-integration-tests/pkg/books"
)

func (s *IntegrationSuite) TestBooksGetBook() {
	client := books.NewBookServiceClient(s.grpcConn)

	s.Run("valid requests", func() {
		tt := []struct {
			name string
			in   *books.GetBookRequest
			out  *books.GetBookResponse
		}{
			{
				name: "book that should be in DB #1",
				in: &books.GetBookRequest{
					Id: 1,
				},
				out: &books.GetBookResponse{
					Book: &books.Book{
						Id:     1,
						Name:   "Martin Eden",
						Author: "Jack London",
					},
				},
			},
			{
				name: "book that should be in DB #2",
				in: &books.GetBookRequest{
					Id: 2,
				},
				out: &books.GetBookResponse{
					Book: &books.Book{
						Id:     2,
						Name:   "Ulysses",
						Author: "James Joyce",
					},
				},
			},
		}

		for _, tc := range tt {
			s.Run(tc.name, func() {
				out, err := client.GetBook(context.Background(), tc.in)
				s.Require().NoError(err)
				s.Require().NotNil(out)

				want := tc.out.Book
				got := out.Book

				s.Assert().Equal(want.Id, got.Id)
				s.Assert().Equal(want.Author, got.Author)
				s.Assert().Equal(want.Name, got.Name)
			})
		}
	})
}

func (s *IntegrationSuite) TestBooksListBooks() {
	client := books.NewBookServiceClient(s.grpcConn)

	s.Run("valid requests", func() {
		tt := []struct {
			name string
			in   *books.ListBooksRequest
			out  *books.ListBooksResponse
		}{
			{
				name: "books that should be in DB",
				in:   &books.ListBooksRequest{},
				out: &books.ListBooksResponse{
					Books: []*books.Book{
						{
							Id:     1,
							Name:   "Martin Eden",
							Author: "Jack London",
						},
						{
							Id:     2,
							Name:   "Ulysses",
							Author: "James Joyce",
						},
					},
				},
			},
		}

		for _, tc := range tt {
			s.Run(tc.name, func() {
				out, err := client.ListBooks(context.Background(), tc.in)
				s.Require().NoError(err)
				s.Require().NotNil(out)

				for i, b := range out.Books {
					s.Assert().Equal(tc.out.Books[i].Id, b.Id)
					s.Assert().Equal(tc.out.Books[i].Name, b.Name)
					s.Assert().Equal(tc.out.Books[i].Author, b.Author)
				}
			})
		}
	})
}
