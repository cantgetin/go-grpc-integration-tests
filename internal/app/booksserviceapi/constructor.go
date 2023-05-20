package booksserviceapi

import (
	"go-integration-tests/internal/database/repository/bookrepository"
	"go-integration-tests/pkg/books"
)

type ServiceImpl struct {
	books.UnimplementedBookServiceServer
	serviceRepository bookrepository.Repository
}

func New(
	serviceRepository bookrepository.Repository,
) *ServiceImpl {
	return &ServiceImpl{
		serviceRepository: serviceRepository,
	}
}
