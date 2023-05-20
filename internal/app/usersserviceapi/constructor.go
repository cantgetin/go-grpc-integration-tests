package userserviceapi

import (
	"go-integration-tests/internal/database/repository/userrepository"
	"go-integration-tests/pkg/users"
)

type ServiceImpl struct {
	users.UnimplementedUserServiceServer
	serviceRepository userrepository.Repository
}

func New(
	serviceRepository userrepository.Repository,
) *ServiceImpl {
	return &ServiceImpl{
		serviceRepository: serviceRepository,
	}
}
